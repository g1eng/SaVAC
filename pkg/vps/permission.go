package vps

import (
	"context"
	"errors"
	"log"
	"net/http"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
)

// ListPermissions lists Permission resources registered on the service.
func (s *SavaClient) ListPermissions() ([]sakuravps.Permission, error) {
	var (
		p    []sakuravps.Permission
		page int32 = 1
	)

	req := s.RawClient.ApiKeyAPI.GetPermissionList(context.Background()).PerPage(100)

	for {
		//log.Printf("page: %d", page)
		res, resp, err := req.Execute()
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			return nil, errors.New("invalid response")
		}
		if res == nil {
			return nil, errors.New("nil returned")
		}
		p = append(p, res.Results...)
		if res.Next.IsSet() && res.Next.Get() != nil {
			log.Printf("page: incremented")
			page += 1
			req = req.Page(page)
		} else {
			break
		}
	}

	return p, nil
}
