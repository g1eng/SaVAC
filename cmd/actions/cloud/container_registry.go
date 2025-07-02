package cloud_actions

import (
	"context"
	"fmt"
	"strconv"

	"github.com/g1eng/savac/cmd/helper"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/search"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/urfave/cli/v3"
)

func getRegistryIdByString(s iaas.ContainerRegistryAPI, arg string) (types.ID, error) {
	id, err := strconv.Atoi(arg)
	if err == nil {
		return types.ID(id), nil
	}
	res, err := s.Find(context.Background(), &iaas.FindCondition{
		Count: 1,
		From:  0,
		Filter: search.Filter{
			search.FilterKey{
				Field: "Name",
				Op:    "=",
			}: arg,
		},
	})
	if err != nil {
		return types.ID(-1), err
	} else if res.Count == 0 {
		return types.ID(-1), fmt.Errorf(" no such container registry found: %s", arg)
	}
	return res.ContainerRegistries[0].ID, nil
}

func findContainerRegistryUIDByUserName(s iaas.ContainerRegistryAPI, id types.ID, name string) (*iaas.ContainerRegistryUser, error) {
	res, err := s.ListUsers(context.Background(), id)
	if err != nil {
		return nil, err
	} else if len(res.Users) == 0 {
		return nil, fmt.Errorf("no container registry users")
	}
	for _, r := range res.Users {
		if r.UserName == name {
			return r, nil
		}
	}
	return nil, fmt.Errorf("no such user: %s", name)
}

func (g *CloudActionGenerator) GenerateContainerRegistryCreateAction(ctx context.Context, cmd *cli.Command) error {
	s := iaas.NewContainerRegistryOp(*g.ApiClient.Caller)
	name := cmd.String("name")
	if cmd.NArg() == 1 {
		if cmd.String("name") != "" {
			return fmt.Errorf("--name can only be used without an argument")
		}
		name = cmd.Args().First()
	}
	req := &iaas.ContainerRegistryCreateRequest{
		Name:           name,
		Description:    cmd.String("description"),
		Tags:           nil,
		AccessLevel:    types.ContainerRegistryAccessLevelMap[cmd.String("permission")],
		SubDomainLabel: name,
	}
	if cmd.IsSet("domain") {
		req.SetVirtualDomain(cmd.String("domain"))
	}
	res, err := s.Create(context.Background(), req)
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}
func (g *CloudActionGenerator) GenerateContainerRegistryDeleteAction(ctx context.Context, cmd *cli.Command) error {
	s := iaas.NewContainerRegistryOp(*g.ApiClient.Caller)
	if cmd.NArg() == 0 {
		return fmt.Errorf("must specify container registry id")
	}
	id, err := getRegistryIdByString(s, cmd.Args().First())
	if err != nil {
		return err
	}
	return s.Delete(context.Background(), id)
}

func (g *CloudActionGenerator) GenerateContainerRegistryListAction(ctx context.Context, cmd *cli.Command) error {
	s := iaas.NewContainerRegistryOp(*g.ApiClient.Caller)
	res, err := s.Find(context.Background(), &iaas.FindCondition{
		Count:   100,
		From:    0,
		Sort:    nil,
		Filter:  nil,
		Include: nil,
		Exclude: nil,
	})
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}

func (g *CloudActionGenerator) GenerateContainerRegistryUserCreateAction(ctx context.Context, cmd *cli.Command) error {
	s := iaas.NewContainerRegistryOp(*g.ApiClient.Caller)
	if cmd.NArg() == 0 {
		return fmt.Errorf("must specify container registry id")
	}
	id, err := getRegistryIdByString(s, cmd.Args().First())
	if err != nil {
		return err
	}
	uName := cmd.String("user")
	password := cmd.String("password")
	permission := cmd.String("permission")
	req := iaas.ContainerRegistryUserCreateRequest{}
	req.SetUserName(uName)
	req.SetPassword(password)
	req.SetPermission(types.EContainerRegistryPermission(permission))
	err = s.AddUser(context.Background(), id, &req)
	if err != nil {
		return err
	}
	return nil
}

func (g *CloudActionGenerator) GenerateContainerRegistryUserListAction(ctx context.Context, cmd *cli.Command) error {
	s := iaas.NewContainerRegistryOp(*g.ApiClient.Caller)
	if cmd.NArg() == 0 {
		return fmt.Errorf("must specify container registry id")
	}
	id, err := getRegistryIdByString(s, cmd.Args().First())
	if err != nil {
		return err
	}
	res, err := s.ListUsers(context.Background(), id)
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}

func (g *CloudActionGenerator) GenerateContainerRegistryUserDeleteAction(ctx context.Context, cmd *cli.Command) error {
	s := iaas.NewContainerRegistryOp(*g.ApiClient.Caller)
	id, err := getRegistryIdByString(s, cmd.Args().First())
	if err != nil {
		return err
	}
	//uid, err := findContainerRegistryUIDByUserName(s, id, cmd.String("user"))
	//uid.
	user := cmd.String("user")
	if user == "" {
		return fmt.Errorf("must specify user name with --user option")
	}
	return s.DeleteUser(context.Background(), id, user)
}
