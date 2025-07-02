package generator

import (
	cloud_actions "github.com/g1eng/savac/cmd/actions/cloud"
	vps_actions "github.com/g1eng/savac/cmd/actions/vps"
	"github.com/g1eng/savac/pkg/cloud/sacloud"
	"github.com/g1eng/savac/pkg/core"
	"github.com/g1eng/savac/pkg/vps"
)

type CloudCommandGenerator struct {
	ApiClient  *sacloud.CloudAPIClient
	OutputType int
	ag         *cloud_actions.CloudActionGenerator
}

func NewCloudCommandGenerator(a *sacloud.CloudAPIClient) *CloudCommandGenerator {
	return &CloudCommandGenerator{
		ApiClient:  a,
		OutputType: core.OutputTypeJson,
		ag: &cloud_actions.CloudActionGenerator{
			ApiClient: a,
		},
	}
}

type VpsCommandGenerator struct {
	ApiClient  *vps.SavaClient
	OutputType int
	NoHeader   bool
	ag         *vps_actions.VpsActionGenerator
}

func NewVpsCommandGenerator(a *vps.SavaClient) *VpsCommandGenerator {
	return &VpsCommandGenerator{
		ApiClient:  a,
		OutputType: core.OutputTypeText,
		ag: &vps_actions.VpsActionGenerator{
			ApiClient: a,
		},
	}
}
