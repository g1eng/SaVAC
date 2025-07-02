package vps

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"github.com/g1eng/savac/pkg/core"
)

// ListRoles lists DummyRoles registered on the service.
func (s *SavaClient) ListRoles() ([]sakuravps.Role, error) {
	var (
		roles []sakuravps.Role
		req         = s.RawClient.ApiKeyAPI.GetRoleList(context.Background()).PerPage(100)
		page  int32 = 1
	)
	for {
		res, rawResp, err := req.Execute()
		if err != nil {
			return nil, err
		}
		msg, _ := io.ReadAll(rawResp.Body)
		if rawResp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("invalid response: %s\n%s", rawResp.Status, msg)
		}
		roles = append(roles, res.Results...)
		if res.Next.IsSet() && res.Next.Get() != nil {
			page++
			req = req.Page(page)
		} else {
			break
		}
	}
	return roles, nil
}

func (s *SavaClient) CreateRole(role *sakuravps.Role) (*sakuravps.Role, error) {
	if role == nil {
		return nil, errors.New("role is nil")
	}
	role, rawResp, err := s.RawClient.ApiKeyAPI.PostRole(context.Background()).Role(*role).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(rawResp, err)
	}
	return role, nil
}

func (s *SavaClient) UpdateRole(roleId int32, rBody *sakuravps.Role) (*sakuravps.Role, error) {
	if rBody == nil {
		return nil, errors.New("role body is nil")
	}
	role, rawResp, err := s.RawClient.ApiKeyAPI.PutRole(context.Background(), roleId).Role(*rBody).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(rawResp, err)
	}
	return role, nil
}

func (s *SavaClient) DeleteRole(id int32) error {
	rawResp, err := s.RawClient.ApiKeyAPI.DeleteRole(context.Background(), id).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}
	return nil
}

func (s *SavaClient) GetRole(id int32) (*sakuravps.Role, error) {
	role, rawResp, err := s.RawClient.ApiKeyAPI.GetRole(context.Background(), id).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(rawResp, err)
	}
	return role, nil
}

func (s *SavaClient) GetRoleByName(name string) (*sakuravps.Role, error) {
	roles, err := s.ListRoles()
	if err != nil {
		return nil, err
	}
	role, err := core.MatchResourceWithName(roles, name)
	if err != nil {
		return nil, err
	}
	return &role, err
}
