package vps_actions

import (
	"fmt"
	"github.com/g1eng/savac/pkg/core"
	"github.com/g1eng/savac/testutil/test_parameters"
	"github.com/urfave/cli/v3"
)

func (v *VpsActionSuite) TestGenerateSwitchCreateAndDeleteAction() {
	tt := []struct {
		name       string
		targetFunc cli.ActionFunc
		args       []string
		expectErr  bool
	}{
		{
			"should fail to create a new switch without zone",
			actionGenerator.GenerateSwitchCreateAction,
			[]string{test_parameters.SampleNewSwitchName},
			true,
		},
		{
			"should create a new switch with zone",
			actionGenerator.GenerateSwitchCreateAction,
			[]string{"--zone", "is1", test_parameters.SampleNewSwitchName},
			false,
		},
		{
			"should delete the switch",
			actionGenerator.GenerateSwitchDeleteAction,
			[]string{test_parameters.SampleNewSwitchName},
			false,
		},
	}

	for _, tc := range tt {
		err := testAction(tc.targetFunc, tc.args...)
		if tc.expectErr {
			if err == nil {
				v.FailNowf("should fail: ", "%s", tc.name)
			}
		} else if err != nil {
			v.FailNowf("fail: ", "%s: %v", tc.name, err)
		}
	}
}
func (v *VpsActionSuite) TestGenerateSwitchNameAction() {
	origName := test_parameters.DummySwitches[0].Name
	id := fmt.Sprintf("%d", test_parameters.DummySwitches[0].Id)
	err = testAction(actionGenerator.GenerateSwitchNameAction, id)
	if err != nil {
		v.Fail("should get name of the switch: " + err.Error())
	}
	err = testAction(actionGenerator.GenerateSwitchNameAction, id, "test-switch-12345")
	if err != nil {
		v.Fail("should rename the switch: " + err.Error())
	}
	err = testAction(actionGenerator.GenerateSwitchNameAction, id, origName)
	if err != nil {
		v.Fail("should rename the switch again: " + err.Error())
	}
}

func (v *VpsActionSuite) TestGenerateSwitchDescriptionAction() {
	origDesc := test_parameters.DummySwitches[0].Description
	id := fmt.Sprintf("%d", test_parameters.DummySwitches[0].Id)
	err = testAction(actionGenerator.GenerateSwitchDescriptionAction, id)
	if err != nil {
		v.Fail("should get description of the switch: " + err.Error())
	}
	err = testAction(actionGenerator.GenerateSwitchDescriptionAction, id, "sample-description")
	if err != nil {
		v.Fail("should change the description of the switch: " + err.Error())
	}
	err = testAction(actionGenerator.GenerateSwitchDescriptionAction, id, origDesc)
	if err != nil {
		v.Fail("should change the description of the switch again: " + err.Error())
	}
	err = testAction(actionGenerator.GenerateSwitchDescriptionAction, test_parameters.DummySwitches[0].Name)
	if err != nil {
		v.Fail("should get description of the switch with name: " + err.Error())
	}
}

func (v *VpsActionSuite) TestGenerateSwitchListAction() {
	for i := 0; i <= core.OutputTypeText; i++ {
		actionGenerator.OutputType = i
		err = testAction(actionGenerator.GenerateSwitchListAction)
		if err != nil {
			v.Failf("fail", "should list switches: %v", err)
		}
	}
}
