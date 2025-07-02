package vps

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/g1eng/savac/testutil/fake_vps"
	"github.com/g1eng/savac/testutil/test_parameters"
	"github.com/stretchr/testify/suite"
)

type SavaClientSuite struct {
	suite.Suite
	serverCli, monCli, nfsCli, swCli, miscCli, apikeyCli *SavaClient
	// API client which returns always connect to fault server
	faultCli  *SavaClient
	isTestAcc bool
}

func (v *SavaClientSuite) SetupSuite() {
	v.isTestAcc = os.Getenv("TESTACC") != ""
	if v.isTestAcc {
		if os.Getenv("SAKURA_VPS_API_TOKEN") == "" {
			v.T().Skip("SAKURA_VPS_API_TOKEN is not set")
		}
		v.serverCli = NewTestClient()
		v.monCli = NewTestClient()
		v.nfsCli = NewTestClient()
		v.swCli = NewTestClient()
		v.miscCli = NewTestClient()
		v.apikeyCli = NewTestClient()
	} else {
		v.serverCli = NewTestClient(test_parameters.FakeServerEndpoint["pkg/server"])
		v.monCli = NewTestClient(test_parameters.FakeServerEndpoint["pkg/monitoring"])
		v.nfsCli = NewTestClient(test_parameters.FakeServerEndpoint["pkg/nfs"])
		v.swCli = NewTestClient(test_parameters.FakeServerEndpoint["pkg/switch"])
		v.miscCli = NewTestClient(test_parameters.FakeServerEndpoint["pkg/misc"])
		v.apikeyCli = NewTestClient(test_parameters.FakeServerEndpoint["pkg/apikey"])
		for name, addr := range test_parameters.FakeServerEndpoint {
			if strings.Index(name, "pkg") == 0 {
				go func() {
					err := fake_vps.StartFakeServer(addr)
					if err != nil {
						panic("failed to start fake server: " + err.Error())
					}
				}()
			}
		}
	}
	v.faultCli = NewTestClient(test_parameters.FaultEndpoint["pkg/server"])
	go func() {
		_, err := fake_vps.StartFaultServer(test_parameters.FaultEndpoint["pkg/server"])
		if err != nil {
			panic("failed to start fault server: %v" + err.Error())
		}
	}()

	time.Sleep(1 * time.Second)
}

func TestSavaClientSuite(t *testing.T) {
	suite.Run(t, new(SavaClientSuite))
}
