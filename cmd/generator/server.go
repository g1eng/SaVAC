package generator

import (
	"github.com/g1eng/savac/cmd/helper"
	"github.com/urfave/cli/v3"
)

func (g *VpsCommandGenerator) GenerateServerSubcommands(isHidden bool) []*cli.Command {
	return []*cli.Command{
		{
			Name:   "list",
			Usage:  "list servers (cached response)",
			Hidden: isHidden,
			Flags: append(patternMatchingFlags,
				&cli.BoolFlag{
					Name:  "l",
					Usage: "list additional information",
					Value: false,
				},
			),
			Action: g.ag.GenerateServerListAction,
		},
		{
			Name:      "info",
			Aliases:   []string{"get", "show"},
			Usage:     "get server detail information",
			Before:    helper.CheckArgsExist,
			Hidden:    isHidden,
			Flags:     patternMatchingFlags,
			ArgsUsage: "<server>",
			Action:    g.ag.GenerateServerInfoAction,
		},
		{
			Name:      "hostname",
			Aliases:   []string{"hn"},
			Usage:     "set/get server name",
			Before:    helper.CheckArgsExist,
			Hidden:    isHidden,
			ArgsUsage: "<serverId>",
			Action:    g.ag.GenerateServerHostnameAction,
		},
		{
			Name:      "description",
			Aliases:   []string{"desc"},
			Usage:     "set/get server description",
			Before:    helper.CheckArgsExist,
			Hidden:    isHidden,
			ArgsUsage: "<serverId>",
			Action:    g.ag.GenerateServerDescriptionAction,
		},
		{
			Name:      "tag",
			Usage:     "set/get server tag (on server description)",
			Hidden:    isHidden,
			ArgsUsage: "<serverId> [key] [value]",
			Action:    g.ag.GenerateServerDescriptionAction,
		},

		{
			Name:      "interfaces",
			Aliases:   []string{"if"},
			Usage:     "list NIC for server",
			Flags:     patternMatchingFlags,
			Hidden:    isHidden,
			ArgsUsage: "<serverId>",
			Action:    g.ag.GenerateServerInterfaceAction,
		},
		{
			Name:    "connect",
			Aliases: []string{"con"},
			Usage:   "connect server NIC to a switch",
			Before:  helper.CheckTwoArgsExist,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "disconnect",
					Aliases: []string{"d"},
				},
			},
			Hidden:    isHidden,
			ArgsUsage: "<NIC-id> <Switch-id>",
			Action:    g.ag.GenerateServerConnectAction,
		},
		{
			Name:      "ptr",
			Usage:     "set/get server's PTR record",
			ArgsUsage: "<server> <hostname>",
			Before:    helper.CheckTwoArgsExist,
			Hidden:    isHidden,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "ipv4",
					Aliases: []string{"v4"},
					Usage:   "set IPv4 PTR",
					Value:   true,
				},
				&cli.BoolFlag{
					Name:    "ipv6",
					Aliases: []string{"v6"},
					Usage:   "set IPv6 PTR",
					Value:   false,
				},
			},
			Action: g.ag.GenerateServerPtrAction,
		},
		{
			Name:    "start",
			Aliases: []string{"boot", "power-on"},
			Before:  helper.CheckArgsExist,
			Usage:   "boot a server",
			Flags:   patternMatchingFlags,
			Hidden:  isHidden,
			Action:  g.ag.GenerateServerStartAction,
		},
		{
			Name:    "stop",
			Aliases: []string{"shutdown", "halt", "power-off"},
			Usage:   "shutdown a server",
			Before:  helper.CheckArgsExist,
			Hidden:  isHidden,
			Flags: append(patternMatchingFlags,
				&cli.BoolFlag{
					Name:  "force",
					Usage: "force power off",
					Value: false,
				},
			),
			Action: g.ag.GenerateServerStopAction,
		},
		{
			Name:    "reboot",
			Aliases: []string{"force-reboot"},
			Usage:   "reboot a server",
			Before:  helper.CheckArgsExist,
			Flags:   patternMatchingFlags,
			Hidden:  isHidden,
			Action:  g.ag.GenerateServerRebootAction,
		},
	}
}
