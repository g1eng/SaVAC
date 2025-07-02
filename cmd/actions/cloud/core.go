package cloud_actions

import "github.com/g1eng/savac/pkg/cloud/sacloud"

type CloudActionGenerator struct {
	ApiClient  *sacloud.CloudAPIClient
	OutputType int
	Debug      bool
}
