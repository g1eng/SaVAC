package generator

import (
	"context"

	"github.com/g1eng/savac/cmd/helper"
	"github.com/urfave/cli/v3"
)

func (cg *CloudCommandGenerator) GenerateDnsCommand() *cli.Command {
	recordFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "id",
			Usage: "id of the DNS appliance",
		},
		&cli.IntFlag{
			Name:  "ttl",
			Usage: "ttl of the DNS record",
			Value: 3600,
		},
	}
	subCmd := []*cli.Command{
		{
			Name:  "create",
			Usage: "create a DNS appliance",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "name",
					Usage:    "name of the DNS appliance",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "description",
					Usage: "description of the DNS appliance",
				},
				&cli.StringSliceFlag{
					Name:  "tags",
					Usage: "tags of the DNS appliance",
				},
			},
			Action: cg.ag.GenerateDnsApplianceCreateAction,
		},
		{
			Name:   "read",
			Usage:  "read a DNS appliance configuration",
			Flags:  patternMatchingFlags,
			Action: cg.ag.GenerateDnsApplianceReadAction,
		},
		{
			Name:  "export",
			Usage: "export the zone file for a DNS appliance",
			Flags: patternMatchingFlags,
			Before: func(ctx context.Context, command *cli.Command) (context.Context, error) {
				if command.IsSet("output-format") {
					if command.String("output-format") != "text" {
						err := command.Set("output-format", "text")
						if err != nil {
							return ctx, err
						}
					}
				}
				return ctx, nil
			},
			Action: cg.ag.GenerateDnsApplianceReadAction,
		},
		{
			Name:  "import",
			Usage: "import a zone for the DNS appliance",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "file",
					Aliases:  []string{"f"},
					Usage:    "zonefile (txt or json) to import to the DNS appliance",
					Required: true,
				},
			},
			Action: cg.ag.GenerateDnsApplianceRecordImportAction,
		},
		{
			Name:    "delete",
			Usage:   "delete a DNS appliance",
			Aliases: []string{"del"},
			Action:  cg.ag.GenerateDnsApplianceDeleteAction,
		},
		{
			Name:   "list",
			Usage:  "list DNS appliances",
			Action: cg.ag.GenerateDnsApplianceListAction,
		},
		{
			Name:    "record",
			Aliases: []string{"rr"},
			Usage:   "operate a DNS appliance record",
			Commands: []*cli.Command{
				{
					Name:      "list",
					Usage:     "list DNS record for an appliance",
					ArgsUsage: "appliance name",
					Before:    helper.CheckArgsExist,
					Action:    cg.ag.GenerateDnsRecordListAction,
				},
				{
					Name:      "add",
					Usage:     "add a record for a DNS appliance",
					Flags:     recordFlags,
					ArgsUsage: "RTYPE RDATA [TTL]",
					Before:    helper.CheckArgsExist,
					Action:    cg.ag.GenerateDnsRecordAddAction,
				},
				{
					Name:    "delete",
					Usage:   "delete a record for a DNS appliance",
					Aliases: []string{"del"},
					Before:  helper.CheckArgsExist,
					Flags: []cli.Flag{
						&cli.IntFlag{
							Name:  "id",
							Usage: "id of the DNS appliance",
						},
						&cli.BoolFlag{
							Name:    "regex",
							Usage:   "regex filtering",
							Aliases: []string{"r"},
						},
					},
					ArgsUsage: "name (of the records) to delete",
					Action:    cg.ag.GenerateDnsRecordDeleteAction,
				},
			},
		},
	}
	return &cli.Command{
		Name:                  "dns",
		Usage:                 "Operate dns appliances on the sakura cloud",
		Hidden:                false,
		Commands:              subCmd,
		EnableShellCompletion: true,
	}
}
