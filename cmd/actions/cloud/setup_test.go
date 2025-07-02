package cloud_actions

import (
	"context"
	"fmt"
	"math/rand/v2"
	"net/netip"
	"os"
	"testing"
	"time"

	"github.com/g1eng/savac/pkg/core"
	"github.com/g1eng/savac/testutil/fake_cloud"
	"github.com/g1eng/savac/testutil/fake_vps"
	"github.com/stretchr/testify/suite"

	"github.com/g1eng/savac/pkg/cloud/sacloud"
	"github.com/urfave/cli/v3"
)

var (
	sacloudCli, err                                 = sacloud.NewCloudApiClient(true)
	isTestAcc                                       = os.Getenv("TESTACC") != ""
	sampleRegisteredDnsApplianceId                  = "113700186673"
	sampleRegisteredDnsZone                         = "ns-testing.s0csec1.org"
	sampleWebAcceleratorSiteId                      = os.Getenv("SAKURACLOUD_WEBACCEL_SITE_ID")
	sampleWebAcceleratorURL                         = "https://" + os.Getenv("SAKURACLOUD_WEBACCEL_DOMAIN") + "/"
	sampleObjectStorageBucketName                   = os.Getenv("SAVAC_TEST_BUCKET_NAME")
	sampleObjectStorageAccessKey                    = os.Getenv(core.OBJECT_STORAGE_API_KEY_ENV_NAME)
	sampleObjectStorageAccessSecret                 = os.Getenv(core.OBJECT_STORAGE_API_SECRET_ENV_NAME)
	sampleUnregisteredDnsARecordName                = "neko-no-nekko2"
	sampleUnregisteredDnsCNAMERecordName            = "neko-no-nekko3rev"
	sampleUnregisteredDnsCNAMERecordValue           = sampleUnregisteredDnsARecordName + "." + sampleRegisteredDnsZone + "."
	sampleUnregisteredDnsMXRecordValue              = "10 " + sampleUnregisteredDnsCNAMERecordValue
	sampleUnregisteredDnsARecordValue               = "127.0.0.1"
	sampleUnregisteredContainerRegistryId           = fmt.Sprintf("hoge-cr-%x\n", rand.Int64()) //nolint
	sampleUnregisteredContainerRegistryUserName     = "sample-user-5"
	sampleUnregisteredContainerRegistryUserPassword = "sample-password-5"

	fakeIaasEndpoint          = "127.0.0.1:20110"
	fakeObjectStorageEndpoint = "127.0.0.1:20113"
	fakeWebAccelEndpoint      = "127.0.0.1:20112"
)

type Args []string
type TestTable []TestTableElement
type TestTableElement struct {
	Args
	ErrorMessage string
	Type         string
}

func withFlags(cmd *cli.Command, flags []cli.Flag) *cli.Command {
	cmd.Flags = append(cmd.Flags, flags...)
	return cmd
}

func testAction(f func(context.Context, *cli.Command) error, args ...string) error {
	cmd := &cli.Command{
		Name: "app",
		Commands: []*cli.Command{
			{
				Name: "cmd",
				Flags: []cli.Flag{
					//shared flags
					&cli.BoolFlag{Name: "l"},
					&cli.StringFlag{Name: "file"},
					//dns flags
					&cli.StringFlag{Name: "id"},
					&cli.StringFlag{Name: "regex"},
					&cli.IntFlag{Name: "ttl", Value: 3600},
					//container registry flags
					&cli.StringFlag{Name: "user"},
					&cli.StringFlag{Name: "password"},
					&cli.StringFlag{Name: "permission", Value: "readwrite"},
					//object storage flags
					&cli.StringSliceFlag{Name: "rw"},
					&cli.StringSliceFlag{Name: "ro"},
					&cli.StringSliceFlag{Name: "wo"},
					&cli.BoolFlag{Name: "recursive"},
					&cli.BoolFlag{Name: "path"},
					&cli.BoolFlag{Name: "key"},
					//webaccel flags
					&cli.StringFlag{Name: "domain-type", Value: "subdomain"},
					&cli.StringFlag{Name: "request-protocol", Value: "http+https"},
					&cli.StringSliceFlag{Name: "cors"},
					&cli.StringFlag{Name: "host-header"},
					&cli.StringFlag{Name: "origin"},
					&cli.StringFlag{Name: "origin-type", Value: "web"},
					&cli.StringFlag{Name: "origin-protocol"},
					&cli.StringFlag{Name: "domain"},
					&cli.StringFlag{Name: "accept-encoding"},
					&cli.StringSliceFlag{Name: "allow"},
					&cli.StringSliceFlag{Name: "deny"},
					&cli.StringFlag{Name: "region", Value: "jp-north-1"},
					&cli.StringFlag{Name: "endpoint", Value: "s3.isk01.sakurastorage.jp"},
					&cli.StringFlag{Name: "bucket"},
					&cli.StringFlag{Name: "access-key"},
					&cli.StringFlag{Name: "access-secret"},
					&cli.StringFlag{Name: "expired", Value: "1min"},
					&cli.IntFlag{Name: "default-cache-ttl"},
					&cli.IntFlag{Name: "month"},
					&cli.IntFlag{Name: "year"},
					&cli.BoolFlag{Name: "vary"},
					&cli.BoolFlag{Name: "next"},
					&cli.BoolFlag{Name: "lets-encrypt", Value: true},
					&cli.BoolFlag{Name: "purge"},
					&cli.BoolFlag{Name: "disable"},
					&cli.BoolFlag{Name: "random"},
					&cli.BoolFlag{Name: "docindex"},
				},
				Action: f,
			},
		},
	}
	return cmd.Run(context.Background(), append([]string{"app", "cmd"}, args...))
}

type CloudActionSuite struct {
	suite.Suite
	Generator CloudActionGenerator
}

func (s *CloudActionSuite) SetupSuite() {
	if isTestAcc {
		if err != nil {
			s.Fail("test suite initialization failure" + err.Error())
		}
		s.Generator = CloudActionGenerator{ApiClient: sacloudCli}
	} else {
		go fake_vps.StartFakeServer(fakeIaasEndpoint) // nolint

		client := fake_cloud.NewFakeApiClient(fakeIaasEndpoint, fakeObjectStorageEndpoint, fakeWebAccelEndpoint)
		engine := fake_cloud.NewFakeEngine("")
		ap, _ := netip.ParseAddrPort(fakeObjectStorageEndpoint)

		sv := fake_cloud.NewObjectStorageFakeServer(int(ap.Port()), engine)
		go func() {
			err := sv.ListenAndServe()
			if err != nil {
				panic(err)
			}
		}()

		time.Sleep(1 * time.Second)
		s.Generator = CloudActionGenerator{ApiClient: &client, Debug: true}
	}
}

func TestSetUp(t *testing.T) {
	suite.Run(t, new(CloudActionSuite))
}
