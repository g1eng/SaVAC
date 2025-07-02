package vps_actions

import (
	"context"
	"errors"
	"fmt"
	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"strconv"

	"github.com/g1eng/savac/cmd/helper"
	"github.com/g1eng/savac/pkg/core"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v3"
)

func (g *VpsActionGenerator) GenerateApiKeyCreateAction(_ context.Context, cmd *cli.Command) error {
	if !cmd.IsSet("role") {
		return errors.New("associated role is not specified")
	}
	name := cmd.Args().First()
	roleName := cmd.String("role")
	roleId, err := strconv.Atoi(roleName)
	if err != nil {
		role, err := g.ApiClient.GetRoleByName(roleName)
		if err != nil {
			return fmt.Errorf("failed to get role: %w", err)
		}
		roleId = int(role.Id)
	}
	key, err := g.ApiClient.CreateApiKey(name, int32(roleId)) //nolint
	if err != nil {
		return fmt.Errorf("failed to create api key: %w", err)
	}
	switch g.OutputType {
	case core.OutputTypeJson:
		return helper.PrintJson(key)
	case core.OutputTypeYaml:
		return helper.PrintYaml(key)
	default:
		println(key.Token)
		return nil
	}
}

func (g *VpsActionGenerator) GenerateApiKeyListAction(_ context.Context, cmd *cli.Command) error {
	keys, err := g.ApiClient.ListApiKeys()
	if err != nil {
		return err
	}
	if cmd.IsSet("regex") {
		keyList, err := core.SearchResourceWithRegex(keys, cmd.Args().First())
		if err != nil {
			return err
		}
		keys = keyList.([]sakuravps.ApiKey)
	} else if cmd.IsSet("search") {
		keyList, err := core.SearchResourceWithName(keys, cmd.Args().First())
		if err != nil {
			return err
		}
		keys = keyList.([]sakuravps.ApiKey)
	}

	switch g.OutputType {
	case core.OutputTypeJson:
		return helper.PrintJson(keys)
	case core.OutputTypeYaml:
		return helper.PrintYaml(keys)
	default:
		if cmd.Bool("l") {
			t := helper.NewList()
			t.SetAlignment(tablewriter.ALIGN_LEFT)
			t.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
			t.SetHeader([]string{"id", "name", "role ID"})
			for _, key := range keys {
				t.Append([]string{fmt.Sprintf("%d", key.Id), key.Name, fmt.Sprintf("%d", int(key.Role))})
			}
			t.Render()
		} else {
			for _, key := range keys {
				println(key.Id)
			}
		}
	}
	return nil
}

func (g *VpsActionGenerator) GenerateApiKeyDeleteAction(_ context.Context, cmd *cli.Command) error {
	s := cmd.Args().First()
	keyId, err := strconv.Atoi(s)
	if err != nil {
		keys, err := g.ApiClient.ListApiKeys()
		if err != nil {
			return fmt.Errorf("failed to get the list api keys: %w", err)
		}
		key, err := core.MatchResourceWithName(keys, s)
		if err != nil {
			return errors.New("no such key: " + s)
		}
		keyId = int(key.Id)
	}
	return g.ApiClient.DeleteApiKeyById(int32(keyId))
}

func (g *VpsActionGenerator) GenerateApiKeyRotateAction(_ context.Context, cmd *cli.Command) error {
	s := cmd.Args().First()
	keyId, err := strconv.Atoi(s)
	if err != nil {
		keys, err := g.ApiClient.ListApiKeys()
		if err != nil {
			return fmt.Errorf("failed to get the list api keys: %w", err)
		}
		key, err := core.MatchResourceWithName(keys, s)
		if err != nil {
			return fmt.Errorf("no such key: %s", s)
		}
		keyId = int(key.Id)
	}
	res, err := g.ApiClient.RotateApiKey(int32(keyId))
	if err != nil {
		return err
	}
	println(res)
	return nil
}
