package sacloud_test

import (
	"os"
	"testing"

	"github.com/g1eng/savac/pkg/cloud/sacloud"
	"github.com/sacloud/iaas-api-go"
)

func Test_NewCloudApiCaller(t *testing.T) {
	c, err := sacloud.NewCloudApiClient()
	if err != nil {
		t.Fatalf("NewCloudApiClient failed: %s", err)
	} else if c == nil {
		t.Fatalf("NewCloudApiClient should return a caller reference")
	}
	_ = os.Unsetenv(iaas.APIAccessTokenEnvKey)
	if _, err = os.ReadFile(os.Getenv("HOME") + "/.usacloud/default/config.json"); err == nil {
		c, err = sacloud.NewCloudApiClient()
		if err != nil {
			t.Fatalf("NewCloudApiClient failed with profile loading: %s", err)
		} else if c == nil {
			t.Fatalf("NewCloudApiClient should return a caller reference with default profile")
		}
	} else {
		_, err = sacloud.NewCloudApiClient()
		if err == nil {
			t.Fatalf("NewCloudApiClient should return an error")
		}
	}
}
