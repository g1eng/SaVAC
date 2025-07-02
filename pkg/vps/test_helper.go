package vps

import (
	"fmt"
	"os"

	sakuravps "github.com/g1eng/sakura_vps_client_go"
)

// NewTestClient is a test helper to initialize SavaClient struct for test suite.
// If a customApiHost argument is given, the test MockEndpoint is evaluated to be hosted locally,
// and API request is performed via http.
func NewTestClient(customApiHost ...string) *SavaClient {
	conf := sakuravps.NewConfiguration()
	conf.DefaultHeader = map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("SAKURA_VPS_API_SECRET")),
		"Content-Type":  "application/json",
	}
	//conf.UserAgent = core.USER_AGENT + "-rc" + core.RELEASE_CANDIDATE_NUMBER
	// for secret masking
	conf.Debug = false
	if os.Getenv("SAVAC_TEST_DEBUG") != "" {
		conf.Debug = true
	}
	if len(customApiHost) > 0 {
		conf.Host = customApiHost[0]
		conf.Scheme = "http"
	}
	//conf.debug = true
	c := NewClient(sakuravps.NewAPIClient(conf))
	c.Debug = true
	return c
}
