package vps

import (
	"context"
	"fmt"
	"io"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"github.com/g1eng/savac/pkg/core"
)

// CreateSwitch creates a switch with specified parameters
func (s *SavaClient) CreateSwitch(name string, desc string, zone string) error {
	req := sakuravps.PostSwitchRequest{
		Name:        name,
		Description: desc,
		ZoneCode:    zone,
	}
	p := s.RawClient.SwitchAPI.PostSwitch(context.Background())
	_, rawResp, err := p.PostSwitchRequest(req).Execute()
	if err != nil {
		return core.EncodeHttpError(rawResp, err)
	}
	return nil
}

// DeleteSwitch deletes switch with ID
func (s *SavaClient) DeleteSwitch(id int32) error {
	resp, err := s.RawClient.SwitchAPI.DeleteSwitch(context.Background(), id).Execute()
	if err != nil {
		return core.EncodeHttpError(resp, err)
	}
	return nil
}

// GetSwitchList gets a list of switch objects from MockEndpoint
func (s *SavaClient) GetSwitchList() ([]sakuravps.Switch, error) {
	var (
		switches []sakuravps.Switch
		page     int32 = 1
	)
	res, resp, err := s.RawClient.SwitchAPI.GetSwitchList(context.Background()).Execute()
	for {
		if err != nil {
			b, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("%s", string(b))
		}
		switches = append(switches, res.Results...)
		if res.Next.Get() != nil {
			page++
			res, resp, err = s.RawClient.SwitchAPI.GetSwitchList(context.Background()).Page(page).Execute()
		} else {
			break
		}
	}
	return switches, nil
}

// PutSwitchName replace switch name with the specified newName
func (s *SavaClient) PutSwitchName(id int32, newName string) error {
	switchDetail, resp, err := s.RawClient.SwitchAPI.GetSwitch(context.Background(), id).Execute()
	if err != nil || switchDetail == nil {
		return core.EncodeHttpError(resp, err)
	}
	switchDetail.Name = newName
	_, resp, err = s.RawClient.SwitchAPI.PutSwitch(context.Background(), id).Switch_(*switchDetail).Execute()
	if err != nil {
		return core.EncodeHttpError(resp, err)
	}
	return nil
}

// PutSwitchDescription replace switch name with the specified newDescription
func (s *SavaClient) PutSwitchDescription(id int32, newDescription string) error {
	switchDetail, resp, err := s.RawClient.SwitchAPI.GetSwitch(context.Background(), id).Execute()
	if err != nil || switchDetail == nil {
		return core.EncodeHttpError(resp, err)
	}
	switchDetail.Description = newDescription
	_, resp, err = s.RawClient.SwitchAPI.PutSwitch(context.Background(), id).Switch_(*switchDetail).Execute()
	if err != nil {
		return core.EncodeHttpError(resp, err)
	}
	return nil
}

func (s *SavaClient) GetSwitchById(id int32) (*sakuravps.Switch, error) {
	res, resp, err := s.RawClient.SwitchAPI.GetSwitch(context.Background(), id).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(resp, err)
	}
	return res, nil
}
