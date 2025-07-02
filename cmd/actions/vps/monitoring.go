package vps_actions

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"github.com/g1eng/savac/cmd/helper"
	"github.com/g1eng/savac/pkg/core"
	"github.com/urfave/cli/v3"
)

func (g *VpsActionGenerator) GenerateServerMonitoringListAction(_ context.Context, cmd *cli.Command) error {
	if cmd.String("server") == "" {
		bulkResp, err := g.ApiClient.GetAllMonitoring()
		if err != nil {
			return err
		}
		switch g.OutputType {
		case core.OutputTypeJson:
			return helper.PrintJson(bulkResp)
		case core.OutputTypeYaml:
			return helper.PrintYaml(bulkResp)
		default:
			helper.PrintMonitoringList(bulkResp)
		}
		return nil
	} else {
		s := cmd.String("server")

		servers, err := g.ApiClient.GetAllServers()
		if err != nil {
			return fmt.Errorf("failed to get server metadata: %v", err)
		}
		var serverIds []int32
		serverId, err := strconv.Atoi(s)
		if err != nil {
			pat, err := regexp.CompilePOSIX(s)
			if err != nil {
				return fmt.Errorf("failed to compile monitoring server id regexp: %v", err)
			}
			for _, sv := range servers {
				if pat.MatchString(sv.Name) {
					serverIds = append(serverIds, sv.Id)
				}
			}
			if len(serverIds) == 0 {
				return fmt.Errorf("server not found")
			}
		} else {
			serverIds = append(serverIds, int32(serverId))
		}
		for i, svId := range serverIds {
			if i != 0 {
				time.Sleep(defaultRequestCoolingTime)
			}
			mon, err := g.ApiClient.GetMonitoringListByServerId(svId)
			if err != nil {
				return err
			}
			switch g.OutputType {
			case core.OutputTypeJson:
				helper.PrintJson(mon) // nolint
			case core.OutputTypeYaml:
				helper.PrintYaml(mon) // nolint
			default:
				helper.PrintMonitoringDetailTable(map[int32][]core.ServerMonitoringMeta{
					svId: mon,
				})
			}
		}
	}
	return nil
}

func (g *VpsActionGenerator) GenerateServerMonitoringInfoAction(_ context.Context, cmd *cli.Command) error {
	if cmd.NArg() < 1 {
		return fmt.Errorf("missing server id")
	}
	serverPat := strings.Split(cmd.Args().Get(0), ",")

	servers, err := g.ApiClient.GetAllServers()
	if err != nil {
		return fmt.Errorf("failed to get server metadata: %v", err)
	}
	var serverIds []int32
	for _, s := range serverPat {
		serverId, err := strconv.Atoi(s)
		if err != nil {
			pat, err := regexp.CompilePOSIX(s)
			if err != nil {
				return fmt.Errorf("failed to compile monitoring server id regexp: %v", err)
			}
			for _, sv := range servers {
				if pat.MatchString(sv.Name) {
					serverIds = append(serverIds, sv.Id)
				}
			}
			if len(serverIds) == 0 {
				return fmt.Errorf("server not found")
			}
		} else {
			serverIds = append(serverIds, int32(serverId))
		}
	}
	for i, svId := range serverIds {
		if i != 0 {
			time.Sleep(defaultRequestCoolingTime)
		}
		mon, err := g.ApiClient.GetMonitoringListByServerId(svId)
		if err != nil {
			return err
		}
		switch g.OutputType {
		case core.OutputTypeJson:
			return helper.PrintJson(mon)
		case core.OutputTypeYaml:
			return helper.PrintYaml(mon)
		default:
			helper.PrintMonitoringDetailTable(map[int32][]core.ServerMonitoringMeta{
				svId: mon,
			})
		}
	}
	return nil
}

func (g *VpsActionGenerator) GenerateServerMonitoringPingAction(_ context.Context, cmd *cli.Command) error {
	serverIds, monitoringName, err := g.parseMonitoringKeys(cmd)
	if err != nil {
		return err
	}
	not, err := parseNotifierArguments(cmd)
	if err != nil {
		return fmt.Errorf("failed to parse notifier arguments: %v", err)
	}
	g.ApiClient.SetMonitoringIntervalMinutes(int32(cmd.Int("monitoring-interval")))
	for i, svId := range serverIds {
		if i != 0 {
			time.Sleep(defaultRequestCoolingTime)
		}
		err = g.ApiClient.AddPingMonitoringForServer(svId, monitoringName, not)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *VpsActionGenerator) GenerateServerMonitoringSshAction(_ context.Context, cmd *cli.Command) error {
	serverIds, monitoringName, err := g.parseMonitoringKeys(cmd)
	if err != nil {
		return err
	}
	not, err := parseNotifierArguments(cmd)
	if err != nil {
		return err
	}
	g.ApiClient.SetMonitoringIntervalMinutes(int32(cmd.Int("monitoring-interval")))

	for i, svId := range serverIds {
		if i != 0 {
			time.Sleep(defaultRequestCoolingTime)
		}
		err = g.ApiClient.AddSshMonitoringForServer(svId, monitoringName, int32(cmd.Int("port")), not)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *VpsActionGenerator) GenerateServerMonitoringTcpAction(_ context.Context, cmd *cli.Command) error {
	serverIds, monitoringName, err := g.parseMonitoringKeys(cmd)
	if err != nil {
		return err
	}
	not, err := parseNotifierArguments(cmd)
	if err != nil {
		return err
	}
	g.ApiClient.SetMonitoringIntervalMinutes(int32(cmd.Int("monitoring-interval")))

	for i, svId := range serverIds {
		if i != 0 {
			time.Sleep(defaultRequestCoolingTime)
		}
		err = g.ApiClient.AddTcpMonitoringForServer(svId, monitoringName, int32(cmd.Int("port")), not)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *VpsActionGenerator) GenerateServerMonitoringSmtpAction(_ context.Context, cmd *cli.Command) error {
	serverIds, monitoringName, err := g.parseMonitoringKeys(cmd)
	if err != nil {
		return err
	}
	not, err := parseNotifierArguments(cmd)
	if err != nil {
		return err
	}
	g.ApiClient.SetMonitoringIntervalMinutes(int32(cmd.Int("monitoring-interval")))

	for i, svId := range serverIds {
		if i != 0 {
			time.Sleep(defaultRequestCoolingTime)
		}
		err = g.ApiClient.AddSmtpMonitoringForServer(svId, monitoringName, int32(cmd.Int("port")), not)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *VpsActionGenerator) GenerateServerMonitoringPop3Action(_ context.Context, cmd *cli.Command) error {
	serverIds, monitoringName, err := g.parseMonitoringKeys(cmd)
	if err != nil {
		return err
	}
	not, err := parseNotifierArguments(cmd)
	if err != nil {
		return err
	}
	g.ApiClient.SetMonitoringIntervalMinutes(int32(cmd.Int("monitoring-interval")))

	for i, svId := range serverIds {
		if i != 0 {
			time.Sleep(defaultRequestCoolingTime)
		}
		err = g.ApiClient.AddPop3MonitoringForServer(svId, monitoringName, int32(cmd.Int("port")), not)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *VpsActionGenerator) GenerateServerMonitoringHttpAction(_ context.Context, cmd *cli.Command) error {
	serverIds, monitoringName, err := g.parseMonitoringKeys(cmd)
	if err != nil {
		return err
	}
	not, err := parseNotifierArguments(cmd)
	if err != nil {
		return err
	}

	if !cmd.IsSet("host") {
		return fmt.Errorf("--host parameter is required")
	}

	for i, svId := range serverIds {
		if i != 0 {
			time.Sleep(defaultRequestCoolingTime)
		}
		param := core.NewHttpMonitoringTarget(int32(cmd.Int("port")), cmd.String("host"), cmd.String("path"))
		if cred := cmd.String("auth"); cred != "" {
			//guarantee for splitting
			u, p := strings.Split(cred, ":")[0], strings.Split(cred, ":")[1]
			param.BasicUserName = sakuravps.NewNullableString(&u)
			param.BasicAuthPassword = sakuravps.NewNullableString(&p)
		}
		err = g.ApiClient.AddHttpMonitoringForServer(svId, monitoringName, param, not)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *VpsActionGenerator) GenerateServerMonitoringHttpsAction(_ context.Context, cmd *cli.Command) error {
	serverIds, monitoringName, err := g.parseMonitoringKeys(cmd)
	if err != nil {
		return err
	}
	not, err := parseNotifierArguments(cmd)
	if err != nil {
		return err
	}
	if !cmd.IsSet("host") {
		return fmt.Errorf("--host parameter is required")
	}
	g.ApiClient.SetMonitoringIntervalMinutes(int32(cmd.Int("monitoring-interval")))

	for i, svId := range serverIds {
		if i != 0 {
			time.Sleep(defaultRequestCoolingTime)
		}
		param := core.NewHttpMonitoringTarget(int32(cmd.Int("port")), cmd.String("host"), cmd.String("path"))
		if cred := cmd.String("auth"); cred != "" {
			//guarantee for splitting
			u, p := strings.Split(cred, ":")[0], strings.Split(cred, ":")[1]
			param.BasicUserName = sakuravps.NewNullableString(&u)
			param.BasicAuthPassword = sakuravps.NewNullableString(&p)
		}
		err = g.ApiClient.AddHttpsMonitoringForServer(svId, monitoringName, param, not)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *VpsActionGenerator) parseMonitoringKeys(cmd *cli.Command) (serverIds []int32, monitoringName string, err error) {
	if cmd.NArg() < 1 {
		return nil, "", fmt.Errorf("missing server id")
	} else if cmd.NArg() < 2 {
		return nil, "", fmt.Errorf("monitoring name must be specified")
	}
	servers, err := g.ApiClient.GetAllServers()
	if err != nil {
		return nil, "", fmt.Errorf("failed to get server metadata: %v", err)
	}
	s := cmd.Args().Get(0)
	serverId, err := strconv.Atoi(s)
	if err != nil {
		//for regex
		serverPat, err := regexp.CompilePOSIX(s)
		if err != nil {
			return nil, "", fmt.Errorf("failed to compile monitoring server id regexp: %v", err)
		}
		for _, sv := range servers {
			if serverPat.MatchString(sv.Name) {
				serverIds = append(serverIds, sv.Id)
			}
		}
	} else {
		serverIds = append(serverIds, int32(serverId))
	}
	return serverIds, cmd.Args().Get(1), nil
}

func parseNotifierArguments(cmd *cli.Command) (*sakuravps.ServerMonitoringSettingsNotification, error) {
	notification := sakuravps.NewServerMonitoringSettingsNotificationWithDefaults()
	arg := cmd.Args().Slice()[2:]
	ni := int32(cmd.Int("notification-interval"))
	mop := cmd.Args()
	log.Println(mop)
	notification.SetIntervalHours(ni)
	if len(arg) == 0 || (len(arg) == 1 && arg[0] == "email") {
		notification.SetEmail(*sakuravps.NewServerMonitoringSettingsNotificationEmail(true))
	} else if len(arg) == 2 && arg[0] == "webhook" {
		//notification.SetEmail(*sakuravps.NewServerMonitoringSettingsNotificationEmail(false))
		notification.SetIncomingWebhook(*sakuravps.NewServerMonitoringSettingsNotificationIncomingWebhook(true, *sakuravps.NewNullableString(&arg[1]), "", ""))
	} else if len(arg) == 4 && arg[0] == "webhook" {
		//notification.SetEmail(*sakuravps.NewServerMonitoringSettingsNotificationEmail(false))
		notification.SetIncomingWebhook(*sakuravps.NewServerMonitoringSettingsNotificationIncomingWebhook(true, *sakuravps.NewNullableString(&arg[1]), arg[2], arg[3]))
	} else {
		return nil, fmt.Errorf("invalid number of arguments: args: %v", arg)
	}
	return notification, nil
}

func (g *VpsActionGenerator) GenerateServerMonitoringDeleteAction(_ context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() < 1 {
		return fmt.Errorf("missing server id")
	}
	var serverIds []int32
	serverPat := cmd.Args().First()
	serverId, err := strconv.Atoi(serverPat)
	if err != nil {
		sv, err := g.ApiClient.GetServerIdsByNamePattern(serverPat)
		if err != nil {
			return err
		}
		serverIds = sv
	} else {
		serverIds = append(serverIds, int32(serverId))
	}
	if cmd.Args().Len() == 1 {
		for _, svId := range serverIds {
			err = g.ApiClient.DeleteAllMonitoringByServerId(svId)
			if err != nil {
				return err
			}
		}
		return nil
	} else if len(serverIds) == 1 {
		m, err := strconv.Atoi(cmd.Args().Get(1))
		if err != nil {
			return fmt.Errorf("invalid monitoring id: %s", cmd.Args().Get(1))
		}
		return g.ApiClient.DeleteMonitoringByServerAndMonitoringId(serverIds[0], int32(m))
	} else {
		return fmt.Errorf("pattern argument cannot accept the additional arguments: %s", cmd.Args().Get(1))
	}
}
