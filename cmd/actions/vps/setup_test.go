package vps_actions

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	fakevps "github.com/g1eng/savac/testutil/fake_vps"
	"github.com/g1eng/savac/testutil/test_parameters"
	"github.com/stretchr/testify/suite"

	"github.com/g1eng/savac/pkg/vps"
	"github.com/urfave/cli/v3"
)

var (
	err                           error
	sampleResolvableServerARecord = os.Getenv("SAKURA_VPS_API_TESTING_HOST")
	sampleWebhookUrl              = os.Getenv("SAVAC_WEBHOOK_URL")
	actionGenerator               *VpsActionGenerator
)

type Args []string
type TestTable []TestTableElement
type TestTableElement struct {
	Args
	ErrorMessage string
	Type         string
}

type VpsActionSuite struct {
	suite.Suite
	apiClient *vps.SavaClient
	isTestAcc bool
}

func (v *VpsActionSuite) SetupSuite() {
	v.isTestAcc = os.Getenv("TESTACC") != ""
	if v.isTestAcc {
		v.apiClient = vps.NewTestClient()
	} else {
		log.Print("custom endpoint: ", test_parameters.FakeServerEndpoint["cmd"])
		v.apiClient = vps.NewTestClient(test_parameters.FakeServerEndpoint["cmd"])
		sampleWebhookUrl = "https://dummy-slack.example.com/some/channel/leak"
	}

	actionGenerator = &VpsActionGenerator{ApiClient: v.apiClient}
	go func() {
		err := fakevps.StartFakeServer(test_parameters.FakeServerEndpoint["cmd"])
		if err != nil {
			v.FailNowf("failed to start fake server", "%v", err)
		}
	}()
	time.Sleep(1 * time.Second)
}

func testAction(f func(ctx context.Context, command *cli.Command) error, args ...string) error {
	cmd := &cli.Command{
		Name: "app",
		Commands: []*cli.Command{
			{
				Name: "cmd",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "l",
					},
					&cli.BoolFlag{
						Name:    "ipv4",
						Aliases: []string{"v4"},
						Usage:   "ipv4 ptr",
					},
					&cli.BoolFlag{
						Name:    "ipv6",
						Aliases: []string{"v6"},
						Usage:   "ipv6 ptr",
					},
					&cli.BoolFlag{
						Name: "disconnect",
					},
					&cli.StringFlag{
						Name: "zone",
					},
					&cli.IntFlag{
						Name:    "notification-interval",
						Aliases: []string{"n"},
						Value:   1,
						Usage:   "notification interval in hours",
					},
					&cli.IntFlag{
						Name:    "monitoring-interval",
						Aliases: []string{"m"},
						Value:   5,
						Usage:   "monitoring interval in minutes",
					},
					&cli.StringFlag{
						Name:  "server",
						Usage: "server target",
					},
					&cli.IntFlag{
						Name:  "port",
						Usage: "destination port",
						Value: 2345,
					},
					&cli.StringFlag{
						Name:  "host",
						Usage: "http host",
					},
					&cli.StringFlag{
						Name:  "path",
						Usage: "http path",
						Value: "/",
					},
					&cli.IntFlag{
						Name:  "status",
						Usage: "http status code",
						Value: 200,
					},
					&cli.BoolFlag{
						Name:  "sni",
						Value: false,
					},
					&cli.StringFlag{
						Name:  "permissions",
						Usage: "permissions list",
					},
					&cli.StringFlag{
						Name:  "switch",
						Usage: "switch resource list for resource filtering",
					},
					&cli.StringFlag{
						Name:  "nfs",
						Usage: "nfs resource list for resource filtering",
					},
					&cli.StringFlag{
						Name: "id",
					},
					&cli.BoolFlag{
						Name:    "regex",
						Aliases: []string{"E"},
					},
					&cli.BoolFlag{
						Name:    "search",
						Aliases: []string{"s"},
					},
					&cli.BoolFlag{
						Name:    "tags",
						Aliases: []string{"T"},
					},
					&cli.StringFlag{
						Name: "role",
					},
					&cli.IntFlag{
						Name:  "ttl",
						Value: 3600,
					},
					&cli.BoolFlag{
						Name:  "force",
						Value: false,
					},
					&cli.BoolFlag{
						Name:    "json",
						Aliases: []string{"j"},
						Value:   false,
					},
					&cli.BoolFlag{
						Name:    "yaml",
						Aliases: []string{"y"},
						Value:   false,
					},
				},
				Action: f,
			},
		},
	}
	return cmd.Run(context.Background(), append([]string{"app", "cmd"}, args...))
}

func TestVpsActionSuite(t *testing.T) {
	suite.Run(t, new(VpsActionSuite))
}
