package sacloud

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/smithy-go/logging"
	"github.com/g1eng/savac/pkg/cloud/model/apprun"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/g1eng/savac/cmd/consts"
	"github.com/g1eng/savac/pkg/cloud/model/object_storage"
	"github.com/g1eng/savac/pkg/core"
	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/api-client-go/profile"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/helper/api"
	"github.com/sacloud/webaccel-api-go"
)

// Config is the usacloud's config object picked from github.com/sacloud/usacloud.
type Config struct {
	profile.ConfigValue

	// Profile プロファイル名
	Profile string `json:"-"`

	// DefaultOutputType デフォルトアウトプットタイプ
	DefaultOutputType string

	// NoColor ANSIエスケープシーケンスによる色つけを無効化
	NoColor bool

	// ProcessTimeoutSec コマンド全体の実行タイムアウトまでの秒数
	ProcessTimeoutSec int

	// ArgumentMatchMode 引数でリソースを特定する際にリソース名と引数を比較する方法を指定
	// 有効な値:
	//   - `partial`(デフォルト): 部分一致
	//   - `exact`: 完全一致
	// Note: 引数はID or Name or Tagsと比較されるが、この項目はNameとの比較にのみ影響する。IDとTagsは常に完全一致となる。
	ArgumentMatchMode string

	// DefaultQueryDriver 各コマンドで--query-driverが省略された場合のデフォルト値
	// 有効な値:
	//   - `jmespath`(デフォルト): JMESPath
	//   - `jq` : gojq
	DefaultQueryDriver string
}

type CloudAPIClient struct {
	Caller           *iaas.APICaller
	WebAccelAPI      *webaccel.Op
	AppRunAPI        *apprun.APIClient
	ObjectStorageAPI *object_storage.APIClient
	S3Client         *s3.Client
	Debug            bool
}

// NewCloudApiClient instantiates a CloudAPIClient object, which is used by Sakura Cloud-related
// command generator in SaVAC.
func NewCloudApiClient(debug ...bool) (*CloudAPIClient, error) {
	var (
		accessToken, accessSecret, zone string
		isDebugMode                     = false
	)
	if len(debug) != 0 {
		isDebugMode = debug[0]
	} else {
		switch os.Getenv("SAVAC_DEBUG") {
		case "1":
			isDebugMode = true
		case "true":
			isDebugMode = true
		}
	}
	if os.Getenv(iaas.APIAccessTokenEnvKey) != "" && os.Getenv(iaas.APIAccessSecretEnvKey) != "" {
		accessToken = os.Getenv(iaas.APIAccessTokenEnvKey)
		accessSecret = os.Getenv(iaas.APIAccessSecretEnvKey)
		zone = os.Getenv(consts.CLUOD_API_DFEAULT_ZONE)
	} else {
		var res Config
		err := profile.Load("default", &res)
		if err != nil {
			return nil, err
		}
		accessToken = res.AccessToken
		accessSecret = res.AccessTokenSecret
		zone = res.DefaultZone
		if res.Zone != "" {
			zone = res.Zone
		}
	}
	if accessToken == "" || accessSecret == "" {
		return nil, fmt.Errorf("no credentials")
	}
	// initialize Sacloud API caller
	iaasCaller := api.NewCallerWithOptions(&api.CallerOptions{
		Options: &client.Options{
			AccessToken:       accessToken,
			AccessTokenSecret: accessSecret,
			AcceptLanguage:    "",
			RetryMax:          3,
			RetryWaitMax:      3,
			RetryWaitMin:      1,
			UserAgent:         core.USER_AGENT,
		},
		APIRootURL:    "",
		DefaultZone:   zone,
		TraceAPI:      isDebugMode,
		FakeMode:      false,
		FakeStorePath: "",
	})

	// initialize Webaccel API caller
	webAccelClient := webaccel.NewOp(&webaccel.Client{
		AccessToken:       accessToken,
		AccessTokenSecret: accessSecret,
		APIRootURL:        "",
		Options: &client.Options{
			AcceptLanguage: "",
			RetryMax:       3,
			RetryWaitMax:   3,
			RetryWaitMin:   1,
			UserAgent:      core.USER_AGENT,
			Trace:          isDebugMode,
		},
	}).(*webaccel.Op)

	// initialize sacloud object storage API client (not S3 compatible client)
	objsCfg := object_storage.NewConfiguration()
	objsCfg.Debug = isDebugMode
	objsCfg.UserAgent = core.USER_AGENT
	basicString := base64.StdEncoding.EncodeToString([]byte(accessToken + ":" + accessSecret))
	objsCfg.DefaultHeader = map[string]string{
		"Authorization": fmt.Sprintf("Basic %s", basicString),
	}
	objsClient := object_storage.NewAPIClient(objsCfg)

	s3Client, err := initializeS3Client()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	appRunCfg := apprun.NewConfiguration()
	appRunCfg.Debug = isDebugMode
	appRunCfg.UserAgent = core.USER_AGENT
	appRunCfg.DefaultHeader = map[string]string{
		"Authorization": fmt.Sprintf("Basic %s", basicString),
	}
	appRunClient := apprun.NewAPIClient(appRunCfg)

	return &CloudAPIClient{
		Caller:           &iaasCaller,
		WebAccelAPI:      webAccelClient,
		ObjectStorageAPI: objsClient,
		S3Client:         s3Client,
		AppRunAPI:        appRunClient,
	}, nil
}

// initializeS3Client initializes S3 compatible client for savac.
func initializeS3Client() (*s3.Client, error) {
	profileName := os.Getenv(core.OBJECT_STORAGE_AWS_PROFILE_ENV_NAME)
	if profileName == "" {
		profileName = core.OBJECT_STORAGE_DEFAULT_PROFILE_NAME
	}
	var (
		cfg      aws.Config
		s3Client *s3.Client
		err      error
	)
	if os.Getenv(core.OBJECT_STORAGE_API_KEY_ENV_NAME) == "" {
		cfg, err = config.LoadDefaultConfig(context.Background(),
			config.WithSharedConfigProfile(profileName),
			config.WithRegion(core.OBJECT_STORAGE_DEFAULT_S3_REGION),
			config.WithBaseEndpoint(core.OBJECT_STORAGE_DEFAULT_S3_ENDPOINT),
		)
	} else {
		cfg, err = config.LoadDefaultConfig(context.Background(),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					os.Getenv(core.OBJECT_STORAGE_API_KEY_ENV_NAME),
					os.Getenv(core.OBJECT_STORAGE_API_SECRET_ENV_NAME),
					""),
			),
			config.WithRegion(core.OBJECT_STORAGE_DEFAULT_S3_REGION),
			config.WithBaseEndpoint(core.OBJECT_STORAGE_DEFAULT_S3_ENDPOINT),
		)
	}
	if err != nil {
		return nil, fmt.Errorf("[WARN] ignoring S3 credentials: unable to load AWS S3 SDK config: %v", err)
	}

	cfg.Logger = logging.NewStandardLogger(os.Stderr)
	cfg.RetryMaxAttempts = 2
	s3Client = s3.NewFromConfig(cfg)
	return s3Client, nil
}
