package generator

import (
	"context"
	"errors"

	"github.com/g1eng/savac/cmd/helper"
	"github.com/urfave/cli/v3"
)

func (g *VpsCommandGenerator) GenerateApiKeySubcommands() *cli.Command {
	return &cli.Command{
		Name:  "apikey",
		Usage: "Operate API key",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Usage:   "list API keys",
				Aliases: []string{"ls"},
				Flags: append([]cli.Flag{
					&cli.BoolFlag{
						Name:  "l",
						Usage: "list additional information",
						Value: false,
					},
				}, patternMatchingFlags...),
				Action: g.ag.GenerateApiKeyListAction,
			},
			{
				Name:    "create",
				Aliases: []string{"cr"},
				Before:  helper.CheckArgsExist,
				Usage:   "create a new API key",
				Hidden:  true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "role",
						Required: true,
					},
				},
				ArgsUsage: "<name> [description]",
				Action:    g.ag.GenerateApiKeyCreateAction,
			},
			{
				Name:      "delete",
				Aliases:   []string{"del"},
				Usage:     "delete the API key",
				Before:    helper.CheckArgsExist,
				Flags:     patternMatchingFlags,
				Hidden:    true,
				ArgsUsage: "<api-key-id>",
				Action:    g.ag.GenerateApiKeyDeleteAction,
			},
			{
				Name:      "rotate",
				Usage:     "rotate the API key",
				Before:    helper.CheckArgsExist,
				Flags:     patternMatchingFlags,
				ArgsUsage: "<api-key-id>",
				Action:    g.ag.GenerateApiKeyRotateAction,
			},
			{
				Name:      "update",
				Usage:     "update the API key information",
				Before:    helper.CheckArgsExist,
				Hidden:    true,
				ArgsUsage: "<api-key-id> <name> <description>",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return errors.New("not implemented")
				},
			},
		},
	}
}
