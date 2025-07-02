package vps_actions

import (
	"context"
	"fmt"
	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"strconv"
	"strings"

	"github.com/g1eng/savac/cmd/helper"
	"github.com/g1eng/savac/pkg/core"
	"github.com/urfave/cli/v3"
)

func (g *VpsActionGenerator) GenerateSwitchListAction(ctx context.Context, cmd *cli.Command) error {
	res, err := g.ApiClient.GetSwitchList()
	if err != nil {
		return err
	}
	if cmd.IsSet("regex") {
		switches, err := core.SearchResourceWithRegex(res, cmd.Args().First())
		if err != nil {
			return err
		}
		res = switches.([]sakuravps.Switch)
	} else if cmd.IsSet("search") {
		switches, err := core.SearchResourceWithName(res, cmd.Args().First())
		if err != nil {
			return err
		}
		res = switches.([]sakuravps.Switch)
	}

	switch g.OutputType {
	case core.OutputTypeJson:
		return helper.PrintJson(res)
	case core.OutputTypeYaml:
		return helper.PrintYaml(res)
	case core.OutputTypeText:
		t := helper.NewList()
		if !g.NoHeader {
			t.SetHeader([]string{"id", "name", "code", "zone", "server_interfaces", "external_connection"})
		}
		for _, d := range res {
			t.Append([]string{
				strconv.Itoa(int(d.Id)),
				d.Name,
				d.SwitchCode,
				d.Zone.Code,
				func() string {
					var interfaces []string
					for _, v := range d.ServerInterfaces {
						v := strconv.Itoa(int(v))
						interfaces = append(interfaces, v)
					}
					for _, v := range d.NfsServerInterfaces {
						v := strconv.Itoa(int(v))
						interfaces = append(interfaces, v)
					}
					return strings.Join(interfaces, ",")
				}(),
				func() string {
					if d.ExternalConnection.IsSet() {
						return d.ExternalConnection.Get().GetServiceCode()
					} else {
						return ""
					}
				}(),
			})
		}
		t.Render()
	}
	return nil
}

func (g *VpsActionGenerator) GenerateSwitchCreateAction(ctx context.Context, cmd *cli.Command) error {
	if !cmd.Args().Present() {
		return fmt.Errorf("no switch name")
	}
	if cmd.String("zone") == "" {
		return fmt.Errorf("zone is required for switch creation")
	}
	s := cmd.Args().First()
	err := g.ApiClient.CreateSwitch(
		s,
		cmd.String("description"),
		cmd.String("zone"),
	)
	return err
}

func (g *VpsActionGenerator) GenerateSwitchDeleteAction(ctx context.Context, cmd *cli.Command) error {
	var (
		swList []sakuravps.Switch
		errs   []error
	)
	s := cmd.Args().First()
	_, err := strconv.Atoi(s)
	if err != nil {
		swList, err = g.ApiClient.GetSwitchList()
		if err != nil {
			return fmt.Errorf("failed to get switch list: %w", err)
		}
		if cmd.IsSet("regex") {
			switches, err := core.SearchResourceWithRegex(swList, cmd.Args().First())
			if err != nil {
				return err
			}
			swList = switches.([]sakuravps.Switch)
		} else if cmd.IsSet("search") {
			switches, err := core.SearchResourceWithName(swList, cmd.Args().First())
			if err != nil {
				return err
			}
			swList = switches.([]sakuravps.Switch)
		} else {
			sw, err := core.MatchResourceWithName(swList, s)
			if err != nil {
				return err
			}
			swList = []sakuravps.Switch{sw}
		}
		for _, sw := range swList {
			err = g.ApiClient.DeleteSwitch(sw.Id)
			if err != nil {
				errs = append(errs, err)
			}
		}
	} else {
		for _, idString := range cmd.Args().Slice() {
			id, err := strconv.Atoi(idString)
			if err != nil {
				errs = append(errs, err)
			} else {
				err = g.ApiClient.DeleteSwitch(int32(id))
				if err != nil {
					errs = append(errs, err)
				}
			}
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("%v", errs)
	}
	return nil
}

func (g *VpsActionGenerator) GenerateSwitchNameAction(ctx context.Context, cmd *cli.Command) error {
	switch cmd.Args().Len() {
	case 1:
		sid, err := strconv.Atoi(cmd.Args().First())
		if err != nil {
			return err
		}
		sw, err := g.ApiClient.GetSwitchById(int32(sid))
		if err != nil {
			return err
		}
		fmt.Println(sw.Name)
		return nil
	case 2:
		sid := cmd.Args().Get(0)
		name := cmd.Args().Get(1)
		i, err := strconv.Atoi(sid)
		if err != nil {
			return err
		}
		return g.ApiClient.PutSwitchName(int32(i), name)
	}
	return fmt.Errorf("at least a switch id must be specified")
}

func (g *VpsActionGenerator) GenerateSwitchDescriptionAction(ctx context.Context, cmd *cli.Command) error {
	s := cmd.Args().First()
	swId, err := strconv.Atoi(s)
	if err != nil {
		swList, err := g.ApiClient.GetSwitchList()
		if err != nil {
			return fmt.Errorf("failed to get switch list: %w", err)
		}
		sw, err := core.MatchResourceWithName(swList, s)
		if err != nil {
			return err
		}
		swId = int(sw.Id)
	}
	switch cmd.Args().Len() {
	case 1:
		sw, err := g.ApiClient.GetSwitchById(int32(swId))
		if err != nil {
			return err
		}
		fmt.Println(sw.Description)
	case 2:
		desc := cmd.Args().Get(1)
		return g.ApiClient.PutSwitchDescription(int32(swId), desc)
	}
	return nil
}
