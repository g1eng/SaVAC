package vps

import (
	"time"

	"github.com/g1eng/savac/testutil/test_parameters"
)

func (v *SavaClientSuite) Test_ApiKeyList() {
	if v.isTestAcc {
		time.Sleep(time.Second)
	}
	keys, err := v.apikeyCli.ListApiKeys()
	v.NoError(err)
	v.NotEmpty(keys)
}

func (v *SavaClientSuite) Test_ApiKeyList_With_FaultServer() {
	if v.isTestAcc {
		time.Sleep(time.Second)
	}
	keys, err := v.faultCli.ListApiKeys()
	if err == nil {
		v.Fail("api key list should fail")
	}
	if len(keys) > 0 {
		v.Fail("api key list should return no payload")
	}
}

func (v *SavaClientSuite) Test_ApiKeyCreate() {
	time.Sleep(time.Second)
	key, err := v.apikeyCli.CreateApiKey(test_parameters.DefaultApiKeyName, test_parameters.SampleRegisteredRoleId)
	if err != nil {
		v.Fail("%v", err)
	}
	test_parameters.SampleRegisteredApiKeyId = key.Id
}

func (v *SavaClientSuite) Test_ApiKeyCreate_With_FaultServer() {
	key, err := v.faultCli.CreateApiKey(test_parameters.DefaultApiKeyName, test_parameters.SampleRegisteredRoleId)
	v.Error(err)
	v.Nil(key)
}

func (v *SavaClientSuite) Test_ApiKeyGet() {
	key, err := v.apikeyCli.GetApiKeyById(test_parameters.SampleRegisteredApiKeyId)
	if err != nil {
		v.Fail("%v", err)
	}
	if key == nil {
		v.Fail("no api key found")
	}
}

func (v *SavaClientSuite) Test_ApiKeyGet_With_FaultServer() {
	if v.isTestAcc {
		time.Sleep(time.Second)
	}
	key, err := v.faultCli.GetApiKeyById(test_parameters.SampleRegisteredApiKeyId)
	if err == nil {
		v.Fail("api key get should fail")
	}
	if key != nil {
		v.Fail("api key get should return nil")
	}
}

func (v *SavaClientSuite) TestSavaClient_RotateApiKey() {
	token, err := v.apikeyCli.RotateApiKey(test_parameters.SampleRegisteredApiKeyId)
	if err != nil {
		v.Fail("%v", err)
	}
	if token == "" {
		v.Fail("no token rotated")
	}
}

func (v *SavaClientSuite) Test_ApiKeyDelete() {
	if v.isTestAcc {
		time.Sleep(time.Second)
	}
	keys, err := v.apikeyCli.ListApiKeys()
	if err != nil {
		v.Fail("%v", err)
	}
	for _, key := range keys {
		if key.Name == test_parameters.DefaultApiKeyName {
			err = v.apikeyCli.DeleteApiKeyById(key.Id)
			if err != nil {
				v.Fail("%v", err)
			}
			return
		}
	}
	v.Fail("no api key found")
}

func (v *SavaClientSuite) Test_ApiKeyDelete_With_FaultServer() {
	c := v.faultCli
	err := c.DeleteApiKeyById(test_parameters.SampleRegisteredApiKeyId)
	if err == nil {
		v.Fail("api key delete should fail")
	}
}
