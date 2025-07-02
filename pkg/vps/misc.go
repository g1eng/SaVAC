package vps

import (
	"context"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
	"github.com/g1eng/savac/pkg/core"
)

func (s *SavaClient) GetAllZone() ([]sakuravps.Zone, error) {
	res, rawResp, err := s.RawClient.ZoneAPI.GetZoneList(context.Background()).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(rawResp, err)
	}
	return res.Results, nil
}

func (s *SavaClient) ListCDROMs() ([]sakuravps.Disc, error) {
	res, rawResp, err := s.RawClient.DiscAPI.GetDiscList(context.Background()).Execute()
	if err != nil {
		return nil, core.EncodeHttpError(rawResp, err)
	}
	return res.Results, nil
}
