package vps

import (
	sakuravps "github.com/g1eng/sakura_vps_client_go"
)

func (s *SavaClient) SetRawApiClient(c *sakuravps.APIClient) {
	s.RawClient = c
}

func (s *SavaClient) SetDebug(debug bool) {
	s.Debug = debug
}

type SavaClient struct {
	Debug bool // debug flag
	//OutputType int  // output type code. See consts above
	Forced                    bool // whether the operation is forced or not
	NoHeader                  bool // whether the table output has header or not
	RawClient                 *sakuravps.APIClient
	monitoringIntervalMinutes int32
}

func NewClient(c *sakuravps.APIClient) *SavaClient {
	return &SavaClient{
		Debug: false,
		//OutputType: core.OutputTypeTable,
		RawClient: c,
	}
}

func SwitchToTestMode(s *SavaClient) *SavaClient {
	testConf := sakuravps.NewConfiguration()
	testConf.Debug = true
	c := sakuravps.NewAPIClient(testConf)
	s.RawClient = c
	return s
}
