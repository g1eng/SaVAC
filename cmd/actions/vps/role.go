package vps_actions

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"github.com/g1eng/savac/cmd/helper"
	"github.com/g1eng/savac/pkg/core"
	"github.com/urfave/cli/v3"
)

func (g *VpsActionGenerator) GenerateRoleListAction(_ context.Context, cmd *cli.Command) error {
	roles, err := g.ApiClient.ListRoles()
	if err != nil {
		return err
	}
	if cmd.IsSet("regex") {
		roleList, err := core.SearchResourceWithRegex(roles, cmd.Args().First())
		if err != nil {
			return err
		}
		roles = roleList.([]sakuravps.Role)
	} else if cmd.IsSet("search") {
		keyList, err := core.SearchResourceWithName(roles, cmd.Args().First())
		if err != nil {
			return err
		}
		roles = keyList.([]sakuravps.Role)
	}

	switch g.OutputType {
	case core.OutputTypeJson:
		return helper.PrintJson(roles)
	case core.OutputTypeYaml:
		return helper.PrintYaml(roles)
	default:
		if cmd.Bool("l") {
			return helper.PrintRolesDetail(roles)
		} else {
			t := helper.NewList()
			t.SetHeader([]string{"id", "name", "description"})
			for _, role := range roles {
				t.Append([]string{fmt.Sprintf("%d", int(role.Id)), role.Name, role.Description})
			}
			t.Render()
		}
		return nil
	}
}

func (g *VpsActionGenerator) GenerateRoleCreateAction(_ context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() == 0 {
		return errors.New("no role name provided")
	}
	name := cmd.Args().First()
	r, err := g.mapRoleFromCliFlags(&sakuravps.Role{}, cmd)
	if err != nil {
		return err
	}
	r.Name = name
	role, err := g.ApiClient.CreateRole(r)
	if err != nil {
		return err
	}
	switch g.OutputType {
	case core.OutputTypeJson:
		return helper.PrintJson(role)
	case core.OutputTypeYaml:
		return helper.PrintYaml(role)
	default:
		return helper.PrintRolesDetail([]sakuravps.Role{*role})
	}
}

func (g *VpsActionGenerator) GenerateRoleUpdateAction(_ context.Context, cmd *cli.Command) error {
	id, err := strconv.Atoi(cmd.Args().First())
	if err != nil {
		role, err := g.ApiClient.GetRoleByName(cmd.Args().First())
		if err != nil {
			return err
		}
		id = int(role.Id)
	}
	origRole, err := g.ApiClient.GetRole(int32(id))
	if err != nil {
		return err
	}
	r, err := g.mapRoleFromCliFlags(origRole, cmd)
	if err != nil {
		return err
	}
	if cmd.IsSet("name") {
		r.Name = cmd.String("name")
	}
	role, err := g.ApiClient.UpdateRole(int32(id), r)
	if err != nil {
		return err
	}
	switch g.OutputType {
	case core.OutputTypeJson:
		return helper.PrintJson(role)
	case core.OutputTypeYaml:
		return helper.PrintYaml(role)
	default:
		return helper.PrintRolesDetail([]sakuravps.Role{*role})
	}
}

func (g *VpsActionGenerator) GenerateRoleReadAction(_ context.Context, cmd *cli.Command) error {
	id, err := strconv.Atoi(cmd.Args().First())
	if err != nil {
		return err
	}
	role, err := g.ApiClient.GetRole(int32(id))
	if err != nil {
		return err
	}
	switch g.OutputType {
	case core.OutputTypeJson:
		return helper.PrintJson(role)
	case core.OutputTypeYaml:
		return helper.PrintYaml(role)
	default:
		return helper.PrintRolesDetail([]sakuravps.Role{*role})
	}
}

func (g *VpsActionGenerator) GenerateRoleDeleteAction(_ context.Context, cmd *cli.Command) error {
	id, err := strconv.Atoi(cmd.Args().First())
	if err != nil {
		role, err := g.ApiClient.GetRoleByName(cmd.Args().First())
		if err != nil {
			return err
		}
		id = int(role.Id)
	}
	return g.ApiClient.DeleteRole(int32(id))
}

func (g *VpsActionGenerator) mapRoleFromCliFlags(role *sakuravps.Role, cmd *cli.Command) (*sakuravps.Role, error) {
	permPattern := cmd.String("permissions")
	hasPermissionFiltering := len(permPattern) > 0

	serverPattern := cmd.String("server")
	switchPattern := cmd.String("switch")
	nfsServerPattern := cmd.String("nfs")
	hasResourceFiltering := len(serverPattern)+len(switchPattern)+len(nfsServerPattern) > 0

	if hasPermissionFiltering {
		log.Println("permission filtering enabled")
		role.PermissionFiltering = "enabled"
		perm, err := g.ApiClient.ListPermissions()
		if err != nil {
			return nil, err
		}
		permList, err := core.SearchResourceWithRegex(perm, permPattern)
		if err != nil {
			return nil, err
		}
		perm = permList.([]sakuravps.Permission)
		for _, p := range perm {
			role.AllowedPermissions = append(role.AllowedPermissions, p.Code)
		}
	} else {
		role.PermissionFiltering = "disabled"
		role.AllowedPermissions = make([]string, 0)
	}

	if hasResourceFiltering {
		role.ResourceFiltering = "enabled"
		var sv, sw, nfs []int32
		if serverPattern != "" {
			allServer, err := g.ApiClient.GetAllServers()
			if err != nil {
				return nil, err
			}
			servers, err := core.SearchResourceWithRegex(allServer, serverPattern)
			if err != nil {
				return nil, err
			}
			for _, s := range servers.([]*sakuravps.Server) {
				sv = append(sv, s.Id)
			}
		}
		if switchPattern != "" {
			allSwitch, err := g.ApiClient.GetSwitchList()
			if err != nil {
				return nil, err
			}
			servers, err := core.SearchResourceWithRegex(allSwitch, switchPattern)
			if err != nil {
				return nil, err
			}
			for _, s := range servers.([]sakuravps.Switch) {
				sw = append(sw, s.Id)
			}
		}
		if nfsServerPattern != "" {
			allNfsServer, err := g.ApiClient.GetAllNFS()
			if err != nil {
				return nil, err
			}
			servers, err := core.SearchResourceWithRegex(allNfsServer, nfsServerPattern)
			if err != nil {
				return nil, err
			}
			for _, s := range servers.([]sakuravps.NfsServer) {
				nfs = append(nfs, s.Id)
			}
		}
		resources := sakuravps.NewNullableRoleAllowedResources(&sakuravps.RoleAllowedResources{
			Servers:    sv,
			Switches:   sw,
			NfsServers: nfs,
		})
		role.AllowedResources = *resources
	} else {
		role.ResourceFiltering = "disabled"
		role.AllowedResources = *sakuravps.NewNullableRoleAllowedResources(&sakuravps.RoleAllowedResources{})
	}

	if cmd.IsSet("description") {
		role.Description = cmd.String("description")
	}
	return role, nil
}
