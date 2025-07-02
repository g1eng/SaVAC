package sacloud_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/g1eng/savac/pkg/cloud/model/object_storage"
	"github.com/g1eng/savac/pkg/cloud/sacloud"
	"github.com/g1eng/savac/pkg/core"
	"github.com/g1eng/savac/testutil/fake_cloud"
	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/iaas-api-go/helper/api"
	"github.com/stretchr/testify/suite"
)

const defaultBucketName = "dummy-bucket"

var (
	sacloudCli, err    = sacloud.NewCloudApiClient(true)
	faultServerPort    = 18383
	fakeServerPort     = 19001
	minioPort          = 19002
	faultEndpoint      = fmt.Sprintf("127.0.0.1:%d", faultServerPort)
	fakeEndpoint       = fmt.Sprintf("127.0.0.1:%d", fakeServerPort)
	localMinioEndpoint = fmt.Sprintf("http://127.0.0.1:%d", minioPort)

	faultServer = http.Server{ // nolint
		Addr: faultEndpoint,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}),
	}
	stubCli = NewFaultApiClient()
	isCI    = os.Getenv("TESTACC") == ""

	defaultAccountKeyId = func() string {
		if isCI {
			return "10000001"
		} else {
			return ""
		}
	}()
	SampleAccountKeyId            = &defaultAccountKeyId
	SampleAccountSecret           = &defaultAccountKeyId
	sampleBucketName              = generateTestingBucketName()
	newBucketName                 = "this-bucket-does-not-exist-at-now"
	samplePermissionName          = "test-savac-perm"
	sampleLocalObjectName         = "stored_obj.txt"
	sampleLocalObjectPath         = "../../../fixtures/" + sampleLocalObjectName
	sampleRemoteObjectParentKey   = "s3://" + sampleBucketName + "/neko/"
	sampleRemoteObjectPath        = sampleRemoteObjectParentKey + sampleLocalObjectName
	sampleRemoteDeepObjectPath    = sampleRemoteObjectParentKey + "nested/" + sampleLocalObjectName
	sampleRemoteAnotherObjectPath = sampleRemoteObjectParentKey + "meshi/nekkachi"
	samplePermissionIdInt         = 8917
	samplePermissionId            = "8917"
	readonlyPermissionId          = "8917"
	readonlyPermissionIdInt       = 8917
	samplePermissionKeyId         = "blank"
	readonlyPermissionKeyId       = "blank"
)

func generateTestingBucketName() string {
	if b := os.Getenv("SAVAC_TEST_BUCKET_NAME"); b != "" {
		return b
	} else {
		return defaultBucketName
	}
}

func NewFaultApiClient() sacloud.CloudAPIClient {
	d := "dummy"
	caller := api.NewCallerWithOptions(&api.CallerOptions{
		Options: &client.Options{
			AccessToken:       d,
			AccessTokenSecret: d,
			UserAgent:         d,
		},
		APIRootURL: faultEndpoint,
		TraceAPI:   true,
	})

	// initialize sacloud object storage API client
	objsCfg := object_storage.NewConfiguration()
	objsCfg.Debug = true
	objsCfg.UserAgent = d
	objsCfg.Host = faultEndpoint
	objsClient := object_storage.NewAPIClient(objsCfg)

	s3Cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				os.Getenv(core.OBJECT_STORAGE_API_KEY_ENV_NAME),
				os.Getenv(core.OBJECT_STORAGE_API_SECRET_ENV_NAME),
				""),
		),
		config.WithBaseEndpoint(faultEndpoint),
		config.WithRegion(d),
	)
	if err != nil {
		panic(err)
	}
	s3Client := s3.NewFromConfig(s3Cfg)
	return sacloud.CloudAPIClient{
		Caller:           &caller,
		ObjectStorageAPI: objsClient,
		S3Client:         s3Client,
		Debug:            true,
	}
}

type ObjectStorageSuite struct {
	suite.Suite
	Client *sacloud.CloudAPIClient
}

func (s *ObjectStorageSuite) SetupSuite() {
	if isCI {
		engine := fake_cloud.NewFakeEngine(defaultAccountKeyId)
		sv := fake_cloud.NewObjectStorageFakeServer(fakeServerPort, engine)
		go func() {
			err := sv.ListenAndServe()
			if err != nil {
				panic(err)
			}
		}()
		time.Sleep(1 * time.Second)
		c := fake_cloud.NewFakeApiClient(fakeEndpoint, fakeEndpoint, localMinioEndpoint)
		s.Client = &c
		_, _ = s.Client.S3Client.CreateBucket(context.Background(), &s3.CreateBucketInput{Bucket: aws.String(sampleBucketName)})
	} else {
		s.Client, err = sacloud.NewCloudApiClient()
		if err != nil {
			s.Fail(err.Error())
		}
	}
}

func TestSetup(t *testing.T) {
	suite.Run(t, new(ObjectStorageSuite))
}
