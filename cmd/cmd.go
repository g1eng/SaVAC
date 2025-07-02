package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/g1eng/savac/cmd/consts"
	"github.com/g1eng/savac/cmd/generator"
	"github.com/g1eng/savac/pkg/cloud/sacloud"
	"github.com/g1eng/savac/pkg/core"
	"github.com/g1eng/savac/pkg/vps"
	"github.com/urfave/cli/v3"
)

func Generate(vpsApiClient *vps.SavaClient) *cli.Command {
	g := generator.NewVpsCommandGenerator(vpsApiClient)
	serverCommands := g.GenerateServerSubcommands(false)
	switchCommand := g.GenerateSwitchCommand()
	monitoringCommands := g.GenerateMonitoringSubcommands()
	roleCommand := g.GenerateRoleSubcommand()
	permCommand := g.GeneratePermissionCommand()
	nfsCommand := g.GenerateNfsCommand()
	discCommand := g.GenerateDiscCommand()
	zoneCommand := g.GenerateZoneCommand()
	apiKeyCommand := g.GenerateApiKeySubcommands()

	cloudApiClient, err := sacloud.NewCloudApiClient()
	if err != nil {
		log.Printf("[WARN] failed to create cloud API client: %v", err)
	}
	cg := generator.NewCloudCommandGenerator(cloudApiClient)
	dnsCommand := cg.GenerateDnsCommand()
	containerRegistryCommand := cg.GenerateContainerRegistryCommand()
	objectStorageCommand := cg.GenerateObjectStorageCommand()
	waCommand := cg.GenerateWebAccelCommand()
	appRunCommand := cg.GenerateAppRunCommand()

	serverRootCommand := &cli.Command{
		Name:  "server",
		Usage: "Operate server resources",
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			for _, c := range cmd.Commands {
				c.Hidden = false
			}
			return ctx, nil
		},
		EnableShellCompletion: true,
	}
	serverRootCommand.Commands = g.BindGlobalFlagsToVpsCommands(serverCommands)

	rootCmd := g.GenerateServerSubcommands(true)
	rootCmd = append(rootCmd, serverRootCommand)
	rootCmd = append(rootCmd, nfsCommand)
	rootCmd = append(rootCmd, switchCommand)
	rootCmd = append(rootCmd, monitoringCommands)
	rootCmd = append(rootCmd, discCommand)
	rootCmd = append(rootCmd, roleCommand)
	rootCmd = append(rootCmd, permCommand)
	rootCmd = append(rootCmd, zoneCommand)
	rootCmd = append(rootCmd, apiKeyCommand)
	rootCmd = g.BindGlobalFlagsToVpsCommands(rootCmd)
	rootCmd = append(rootCmd, cg.BindGlobalFlagsToCloudCommands([]*cli.Command{
		dnsCommand,
		containerRegistryCommand,
		objectStorageCommand,
		waCommand,
		appRunCommand,
	})...)

	app := &cli.Command{
		Name:    consts.APP_NAME,
		Version: core.VERSION,
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			vpsApiClient.RawClient.GetConfig().DefaultHeader["Authorization"] = fmt.Sprintf("Bearer %s", cmd.String("api-token"))
			vpsApiClient.Debug = cmd.Bool("debug")
			if vpsApiClient.Debug {
				err := cmd.Set("api-token", "dummy-local-api-token")
				if err != nil {
					panic(err)
				}
			}
			if cmd.Bool("test-mode") {
				vps.SwitchToTestMode(vpsApiClient)
			}
			return ctx, nil
		},
		Usage:                 "Sakura VPS Administrator's CLI",
		EnableShellCompletion: true,
	}
	switch regexp.MustCompilePOSIX("^.*/").ReplaceAllString(os.Args[0], "") {
	case consts.APP_NAME:
		app.Commands = rootCmd
	case consts.APP_ALIAS_OBJECT_STORAGE:
		app.Name = consts.APP_ALIAS_OBJECT_STORAGE
		app.Commands = objectStorageCommand.Commands
	case consts.APP_ALIAS_CONTAINER_REGISTRY:
		app.Name = consts.APP_ALIAS_CONTAINER_REGISTRY
		app.Commands = containerRegistryCommand.Commands
	case consts.APP_ALIAS_DNS:
		app.Name = consts.APP_ALIAS_DNS
		app.Commands = dnsCommand.Commands
	default:
		app.Commands = rootCmd
	}
	return g.BindGlobalFlagsToApp(app)
}
