package generator

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/g1eng/savac/cmd/consts"
	"github.com/g1eng/savac/cmd/helper"
	"github.com/g1eng/savac/pkg/cloud/sacloud"
	"github.com/g1eng/savac/pkg/core"
	"github.com/urfave/cli/v3"
)

func (g *VpsCommandGenerator) generateGlobalVPSFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			DefaultText: "false",
			Value:       false,
			Action: func(_ context.Context, command *cli.Command, b bool) error {
				g.ApiClient.RawClient.GetConfig().Debug = b
				g.ag.ApiClient.RawClient.GetConfig().Debug = b
				return nil
			},
		},
		&cli.BoolFlag{
			Name:        "dry-run",
			DefaultText: "false",
			Value:       false,
		},
		&cli.StringFlag{
			Name:    "api-token",
			Sources: cli.EnvVars(consts.API_TOKEN_ENV_NAME),
			Hidden:  true,
		},
		&cli.BoolFlag{
			Name:   "test-mode",
			Hidden: true,
		},
		&cli.StringFlag{
			Name:     "zone",
			Required: false,
		},
		&cli.StringFlag{
			Name:    "output-format",
			Usage:   "[text|json|yaml]",
			Sources: cli.EnvVars(consts.OUTPUT_TYPE_ENV_NAME),
			Action: func(ctx context.Context, cmd *cli.Command, s string) error {
				switch s {
				case "json":
					g.OutputType = core.OutputTypeJson
					g.ag.OutputType = core.OutputTypeJson
				case "yaml":
					g.OutputType = core.OutputTypeYaml
					g.ag.OutputType = core.OutputTypeYaml
				case "table":
					fallthrough
				case "text":
					g.OutputType = core.OutputTypeText
					g.ag.OutputType = core.OutputTypeText
				default:
					return fmt.Errorf("no such output type: %s", s)
				}
				return nil
			},
		},
		&cli.BoolFlag{
			Name:    "json",
			Usage:   "alias for --output-format=json",
			Aliases: []string{"j"},
			Action: func(ctx context.Context, cmd *cli.Command, b bool) error {
				if cmd.Bool("yaml") || cmd.Bool("table") {
					return errors.New("table and yaml cannot be used at the same time")
				}
				if b {
					g.OutputType = core.OutputTypeJson
					g.ag.OutputType = core.OutputTypeJson
				}
				return nil
			},
			Value: false,
		},
		&cli.BoolFlag{
			Name:    "yaml",
			Aliases: []string{"Y"},
			Usage:   "alias for --output-format=yaml",
			Value:   false,
			Action: func(ctx context.Context, cmd *cli.Command, b bool) error {
				if cmd.Bool("table") || cmd.Bool("json") {
					return errors.New("table and json cannot be used at the same time")
				}
				if b {
					g.OutputType = core.OutputTypeYaml
					g.ag.OutputType = core.OutputTypeYaml
				}
				return nil
			},
		},
		&cli.BoolFlag{
			Name:    "text",
			Aliases: []string{"t"},
			Usage:   "alias for --output-format=table",
			Value:   false,
			Action: func(ctx context.Context, cmd *cli.Command, b bool) error {
				if cmd.Bool("yaml") || cmd.Bool("json") {
					return errors.New("yaml and json cannot be used with `--table` at the same time")
				}
				if b {
					g.OutputType = core.OutputTypeText
					g.ag.OutputType = core.OutputTypeText
				}
				return nil
			},
		},
		&cli.BoolFlag{
			Name:    "no-header",
			Aliases: []string{"H"},
			Usage:   "suppress the header output of the table",
			Value:   false,
			Action: func(ctx context.Context, cmd *cli.Command, b bool) error {
				if cmd.Bool("yaml") || cmd.Bool("json") {
					return errors.New("yaml and json cannot be used with the `--no-header` at the same time")
				}
				if b {
					g.OutputType = core.OutputTypeText
					g.ag.OutputType = core.OutputTypeText
					g.NoHeader = true
				}
				return nil
			},
		},
	}
}

func (cg *CloudCommandGenerator) generateGlobalCloudAPIFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			DefaultText: "false",
			Value:       false,
			Action: func(_ context.Context, command *cli.Command, b bool) (err error) {
				cg.ApiClient, err = sacloud.NewCloudApiClient(b)
				if err != nil {
					log.Printf("[WARN] could not create cloud API client: %v", err)
				}
				cg.ag.ApiClient, err = sacloud.NewCloudApiClient(b)
				if err != nil {
					log.Printf("[WARN] could not create cloud API client: %v", err)
				}
				return nil
			},
		},
		&cli.BoolFlag{
			Name:        "dry-run",
			DefaultText: "false",
			Value:       false,
		},
		&cli.StringFlag{
			Name:    "api-token",
			Sources: cli.EnvVars(consts.API_TOKEN_ENV_NAME),
			Hidden:  true,
		},
		&cli.StringFlag{
			Name:     "zone",
			Required: false,
		},
		&cli.StringFlag{
			Name:    "output-format",
			Usage:   "[table|json|yaml]",
			Sources: cli.EnvVars(consts.OUTPUT_TYPE_ENV_NAME),
			Action: func(ctx context.Context, cmd *cli.Command, s string) (err error) {
				cg.OutputType, err = helper.GetOutputDigitByName(s)
				cg.ag.OutputType = cg.OutputType
				return err
			},
		},
		&cli.BoolFlag{
			Name:    "json",
			Aliases: []string{"j"},
			Action: func(ctx context.Context, cmd *cli.Command, b bool) error {
				if cmd.Bool("yaml") || cmd.Bool("table") {
					return errors.New("table and yaml cannot be used at the same time")
				}
				if b {
					cg.OutputType = core.OutputTypeJson
					cg.ag.OutputType = core.OutputTypeJson
				}
				return nil
			},
			Value: false,
		},
		&cli.BoolFlag{
			Name:    "yaml",
			Aliases: []string{"Y"},
			Value:   false,
			Action: func(ctx context.Context, cmd *cli.Command, b bool) error {
				if cmd.Bool("table") || cmd.Bool("json") {
					return errors.New("table and json cannot be used at the same time")
				}
				if b {
					cg.OutputType = core.OutputTypeYaml
					cg.ag.OutputType = core.OutputTypeYaml
				}
				return nil
			},
		},
		&cli.BoolFlag{
			Name:    "txt",
			Aliases: []string{"t"},
			Value:   false,
			Action: func(ctx context.Context, cmd *cli.Command, b bool) error {
				if cmd.Bool("yaml") || cmd.Bool("json") {
					return errors.New("yaml and json cannot be used with `--text` at the same time")
				}
				if b {
					cg.OutputType = core.OutputTypeText
					cg.ag.OutputType = core.OutputTypeText
				}
				return nil
			},
		},
	}
}

func (g *VpsCommandGenerator) BindGlobalFlagsToApp(app *cli.Command) *cli.Command {
	app.Flags = append(app.Flags, g.generateGlobalVPSFlags()...)
	return app
}

func (cg *CloudCommandGenerator) BindGlobalFlagsToApp(app *cli.Command) *cli.Command {
	app.Flags = append(app.Flags, cg.generateGlobalCloudAPIFlags()...)
	return app
}

func (g *VpsCommandGenerator) BindGlobalFlagsToVpsCommands(cmd []*cli.Command) []*cli.Command {
	options := g.generateGlobalVPSFlags()
	for i := range cmd {
		cmd[i].Flags = append(cmd[i].Flags, options...)
		if cmd[i].Commands != nil {
			cmd[i].Commands = g.BindGlobalFlagsToVpsCommands(cmd[i].Commands)
		}
	}
	return cmd
}

func (cg *CloudCommandGenerator) BindGlobalFlagsToCloudCommands(cmd []*cli.Command) []*cli.Command {
	options := cg.generateGlobalCloudAPIFlags()
	for i := range cmd {
		cmd[i].Flags = append(cmd[i].Flags, options...)
		if cmd[i].Commands != nil {
			cmd[i].Commands = cg.BindGlobalFlagsToCloudCommands(cmd[i].Commands)
		}
	}
	return cmd
}
