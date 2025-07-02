package fake_cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/g1eng/savac/pkg/cloud/model/object_storage"
	"github.com/g1eng/savac/pkg/cloud/sacloud"
	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/iaas-api-go/helper/api"
)

func NewFakeApiClient(iaasEndpoint, objectStorageEndpoint, minioEndpoint string) sacloud.CloudAPIClient {
	d := "dummy"
	caller := api.NewCallerWithOptions(&api.CallerOptions{
		Options: &client.Options{
			AccessToken:       d,
			AccessTokenSecret: d,
			UserAgent:         d,
		},
		APIRootURL:    iaasEndpoint,
		FakeMode:      true,
		FakeStorePath: "/",
		TraceAPI:      true,
	})

	// initialize sacloud object storage API client
	objsCfg := object_storage.NewConfiguration()
	objsCfg.Debug = true
	objsCfg.UserAgent = d
	objsCfg.Host = objectStorageEndpoint
	objsCfg.Scheme = "http"
	objsClient := object_storage.NewAPIClient(objsCfg)

	s3Cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				"minioadmin",
				"minioadmin",
				""),
		),
		config.WithBaseEndpoint(minioEndpoint),
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
