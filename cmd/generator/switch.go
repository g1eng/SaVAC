package generator

import (
	"github.com/g1eng/savac/cmd/helper"
	"github.com/urfave/cli/v3"
)

func (g *VpsCommandGenerator) GenerateSwitchCommand() *cli.Command {
	return &cli.Command{
		Name:    "switch",
		Usage:   "Operate switch resources",
		Aliases: []string{"sw"},
		Commands: []*cli.Command{
			{
				Name:    "list",
				Usage:   "list switches",
				Aliases: []string{"ls"},
				Flags:   patternMatchingFlags,
				Action:  g.ag.GenerateSwitchListAction,
			},
			{
				Name:      "create",
				ArgsUsage: "NAME",
				Usage:     "create a new switch",
				Before:    helper.CheckArgsExist,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "description",
						Aliases: []string{"desc"},
					},
				},
				Action: g.ag.GenerateSwitchCreateAction,
			},
			{
				Name:      "delete",
				ArgsUsage: "switchId",
				Usage:     "delete a switch",
				Before:    helper.CheckArgsExist,
				Flags:     patternMatchingFlags,
				Action:    g.ag.GenerateSwitchDeleteAction,
			},
			{
				Name:      "name",
				ArgsUsage: "switchId [name]",
				Before:    helper.CheckArgsExist,
				Usage:     "show or update switch name",
				Action:    g.ag.GenerateSwitchNameAction,
			},
			{
				Name:      "description",
				ArgsUsage: "switchId [descriptions...]",
				Before:    helper.CheckArgsExist,
				Usage:     "show or update switch description",
				Action:    g.ag.GenerateSwitchDescriptionAction,
			},
		},
	}
}
