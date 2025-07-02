package vps_actions

import (
	"fmt"
	"github.com/g1eng/savac/testutil/test_parameters"
	"github.com/urfave/cli/v3"
	"time"
)

func (v *VpsActionSuite) TestGenerateApiKeyAction() {
	//go fake_vps.StartFakeServer(localTestEndpoint)
	if v.isTestAcc {
		time.Sleep(2 * time.Second)
	}
	tt := []struct {
		name string
		args []string
	}{
		{
			"should list api keys",
			[]string{},
		},
		{
			"should dump apikey list with json",
			[]string{"-j"},
		},
		{
			"should dump apikey list with yaml",
			[]string{"-y"},
		},
	}
	for _, tc := range tt {
		err := testAction(actionGenerator.GenerateApiKeyListAction, tc.args...)
		if err != nil {
			v.Failf("testcase: %s\nfailed: %s", tc.name, err.Error())
		}
	}
}

func (v *VpsActionSuite) TestScenarioApiKeyCreateRotateDeleteAction() {
	if v.isTestAcc {
		time.Sleep(15 * time.Second)
	}
	var err error
	tt := []struct {
		name        string
		targetFunc  cli.ActionFunc
		args        []string
		expectError bool
	}{
		{
			"apikey creation without --role option should fail",
			actionGenerator.GenerateApiKeyCreateAction,
			[]string{test_parameters.SampleUnregisteredApiKeyName},
			true,
		},
		{
			"apikey creation with valid role specification should have success",
			actionGenerator.GenerateApiKeyCreateAction,
			[]string{"--role", fmt.Sprintf("%d", test_parameters.SampleRegisteredRoleId), test_parameters.SampleUnregisteredApiKeyName},
			false,
		},
		{
			"apikey rotation for new key with name should have success",
			actionGenerator.GenerateApiKeyRotateAction,
			[]string{test_parameters.SampleUnregisteredApiKeyName},
			false,
		},
		{
			"apikey deletion for the new key with its name should have success",
			actionGenerator.GenerateApiKeyDeleteAction,
			[]string{test_parameters.SampleUnregisteredApiKeyName},
			false,
		},
	}

	for _, tc := range tt {
		err = testAction(tc.targetFunc, tc.args...)
		if tc.expectError {
			if err == nil {
				v.FailNowf("an error expected", "%s", tc.name)
			}
		} else if err != nil {
			v.FailNowf("error", ": %s: %s", tc.name, err.Error())
		}
	}
}
