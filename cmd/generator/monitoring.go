package generator

import (
	"context"
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"
)

// TODO: implement pattern matching for servers

func (g *VpsCommandGenerator) GenerateMonitoringSubcommands() *cli.Command {
	return &cli.Command{
		Name:    "monitoring",
		Usage:   "Control server monitoring",
		Aliases: []string{"mon"},
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List all monitoring",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "protocol",
						Aliases: []string{"P"},
						Usage:   "filter result by monitoring type [ping|tcp|smtp|pop3|http|https]",
						Value:   "",
					},
					&cli.StringFlag{
						Name:  "server",
						Usage: "filter result by monitored server",
					},
				},
				Action: g.ag.GenerateServerMonitoringListAction,
			},
			{
				Name:   "info",
				Usage:  "Show monitoring detail for specific server(s)",
				Flags:  []cli.Flag{},
				Action: g.ag.GenerateServerMonitoringInfoAction,
			},

			{
				Name:      "ping",
				Usage:     "Add ping monitoring to specified server(s)",
				ArgsUsage: "serverId name [email|webhook] [[url] [slack_team] [slack_channel]]",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "notification-interval", Aliases: []string{"n"}, Value: 1, Usage: "notification interval for monitoring in hours"},
					&cli.IntFlag{Name: "monitoring-interval", Aliases: []string{"m"}, Value: 5, Usage: "monitoring interval in minutes"},
				},
				Action: g.ag.GenerateServerMonitoringPingAction,
			},
			{
				Name:      "ssh",
				Usage:     "Add SSH monitoring to specified server(s)",
				ArgsUsage: "serverId name [email|webhook] [[url] [slack_team] [slack_channel]]",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Value: 22, Usage: "destination port for the monitoring"},
					&cli.IntFlag{Name: "notification-interval", Aliases: []string{"n"}, Value: 1, Usage: "notification interval for monitoring in hours"},
					&cli.IntFlag{Name: "monitoring-interval", Aliases: []string{"m"}, Value: 5, Usage: "monitoring interval in minutes"},
				},
				Action: g.ag.GenerateServerMonitoringSshAction,
			},
			{
				Name:      "tcp",
				Usage:     "Add TCP monitoring to specified server(s)",
				ArgsUsage: "serverId name [email|webhook] [[url] [slack_team] [slack_channel]]",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Required: true, Usage: "destination port for the monitoring"},
					&cli.IntFlag{Name: "notification-interval", Aliases: []string{"n"}, Value: 1, Usage: "notification interval for monitoring in hours"},
					&cli.IntFlag{Name: "monitoring-interval", Aliases: []string{"m"}, Value: 5, Usage: "monitoring interval in minutes"},
				},
				Action: g.ag.GenerateServerMonitoringTcpAction,
			},
			{
				Name:  "http",
				Usage: "Add HTTP monitoring",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "notification-interval", Aliases: []string{"n"}, Value: 1, Usage: "notification interval in hours"},
					&cli.IntFlag{Name: "monitoring-interval", Aliases: []string{"m"}, Value: 5, Usage: "monitoring interval in minutes"},
					&cli.StringFlag{Name: "host", Usage: "host header", Required: true},
					&cli.StringFlag{Name: "path", Usage: "monitoring path", Value: "/", Required: false},
					&cli.IntFlag{Name: "port", Usage: "destination port for monitoring", Aliases: []string{"p"}, Value: 80, Required: false},
					&cli.StringFlag{
						Name:     "auth",
						Aliases:  []string{"A"},
						Usage:    "basic authorization credentials in user:password format",
						Required: false,
						Action: func(context context.Context, cmd *cli.Command, s string) error {
							if s == "" {
								return nil
							} else if len(strings.Split(s, ":")) == 1 {
								return fmt.Errorf("invalid credential (the `:` not detected): %s", s)
							}
							return nil
						},
					},
					&cli.IntFlag{Name: "status", Usage: "expected status code", Value: 200, Required: false},
				},
				ArgsUsage: "<serverId> <name> [email|webhook] [url] [slack_team] [slack_channel]",
				Action:    g.ag.GenerateServerMonitoringHttpAction,
			},
			{
				Name:  "https",
				Usage: "Add HTTPS monitoring",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "notification-interval", Aliases: []string{"n"}, Value: 1, Usage: "notification interval in hours"},
					&cli.IntFlag{Name: "monitoring-interval", Aliases: []string{"m"}, Value: 5, Usage: "monitoring interval in minutes"},
					&cli.BoolFlag{Name: "sni", Usage: "the server sets SNI or not", Value: false, Required: false},
					&cli.StringFlag{Name: "host", Usage: "host header", Required: true},
					&cli.StringFlag{Name: "path", Usage: "monitoring path", Value: "/", Required: true},
					&cli.IntFlag{Name: "port", Usage: "destination port for monitoring", Aliases: []string{"p"}, Value: 443, Required: false},
					&cli.StringFlag{
						Name:     "auth",
						Aliases:  []string{"A"},
						Usage:    "basic authorization credentials in user:password format",
						Required: false,
						Action: func(context context.Context, cmd *cli.Command, s string) error {
							if s == "" {
								return nil
							} else if len(strings.Split(s, ":")) == 1 {
								return fmt.Errorf("invalid credential (the `:` not detected): %s", s)
							}
							return nil
						},
					},
					&cli.IntFlag{Name: "status", Usage: "expected status code", Value: 200, Required: false},
				},
				ArgsUsage: "<serverId> <name> [email|webhook] [url] [slack_team] [slack_channel]",
				Action:    g.ag.GenerateServerMonitoringHttpsAction,
			},
			{
				Name:      "smtp",
				Usage:     "Add SMTP monitoring to specified server(s)",
				ArgsUsage: "serverId name [email|webhook] [[url] [slack_team] [slack_channel]]",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Value: 587, Usage: "destination port for the monitoring"},
					&cli.IntFlag{Name: "notification-interval", Aliases: []string{"n"}, Value: 1, Usage: "notification interval for the monitoring in hours"},
					&cli.IntFlag{Name: "monitoring-interval", Aliases: []string{"m"}, Value: 5, Usage: "monitoring interval in minutes"},
				},
				Action: g.ag.GenerateServerMonitoringSmtpAction,
			},
			{
				Name:      "pop3",
				Usage:     "Add POP3 monitoring to specified server(s)",
				ArgsUsage: "serverId name [email|webhook] [[url] [slack_team] [slack_channel]]",
				Flags: []cli.Flag{&cli.IntFlag{Name: "port", Aliases: []string{"p"}, Value: 990, Usage: "destination port for the monitoring"},
					&cli.IntFlag{Name: "notification-interval", Aliases: []string{"n"}, Value: 1, Usage: "notification interval for the monitoring in hours"},
					&cli.IntFlag{Name: "monitoring-interval", Aliases: []string{"m"}, Value: 5, Usage: "monitoring interval in minutes"},
				},
				Action: g.ag.GenerateServerMonitoringPop3Action,
			},
			{
				Name:      "delete",
				Aliases:   []string{"del"},
				ArgsUsage: "<serverId> [monitoringId]",
				Usage:     "delete monitoring from specified server (optionally, with monitoring id)",
				Action:    g.ag.GenerateServerMonitoringDeleteAction,
			},
		},
	}
}
