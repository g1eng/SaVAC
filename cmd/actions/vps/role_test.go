package vps_actions

import (
	"fmt"
	"github.com/g1eng/savac/pkg/core"
	"github.com/g1eng/savac/testutil/test_parameters"
	"time"
)

func (v *VpsActionSuite) TestGenerateRoleListAction() {
	for _, o := range core.OutputTypes {
		if v.isTestAcc {
			time.Sleep(5 * time.Second)
		}
		actionGenerator.OutputType = o
		err := testAction(actionGenerator.GenerateRoleListAction)
		if err != nil {
			v.Failf("fail", "%v", err)
		}
		err = testAction(actionGenerator.GenerateRoleListAction, "-l")
		if err != nil {
			v.Failf("fail", "%v", err)
		}
	}
}

func (v *VpsActionSuite) TestGenerateRoleReadAction() {
	for _, o := range core.OutputTypes {
		if v.isTestAcc {
			time.Sleep(5 * time.Second)
		}
		actionGenerator.OutputType = o
		err := testAction(actionGenerator.GenerateRoleReadAction, fmt.Sprintf("%d", test_parameters.SampleRegisteredRoleId))
		if err != nil {
			v.Failf("fail", "%v", err)
		}
	}
}

func (v *VpsActionSuite) TestScenarioRoleCreateUpdateDeleteAction() {
	if v.isTestAcc {
		time.Sleep(5 * time.Second)
	}
	err := testAction(
		actionGenerator.GenerateRoleCreateAction,
		"--permissions", "put-server",
		"--server", test_parameters.SampleRegisteredServerIdString,
		"--switch", test_parameters.SampleRegisteredNfsSwitchIdText,
		"--nfs", test_parameters.SampleRegisteredNfsIdText,
		test_parameters.SampleUnregisteredRoleName,
	)
	if err != nil {
		v.Failf("fail", "%v", err)
	}
	if v.isTestAcc {
		time.Sleep(5 * time.Second)
	}
	err = testAction(
		actionGenerator.GenerateRoleUpdateAction,
		"--permissions", "put-switch",
		"--switch", test_parameters.SampleRegisteredNfsSwitchIdText,
		test_parameters.SampleUnregisteredRoleName,
	)
	if err != nil {
		v.Failf("fail", "%v", err)
	}
	if v.isTestAcc {
		time.Sleep(5 * time.Second)
	}
	err = testAction(actionGenerator.GenerateRoleDeleteAction, test_parameters.SampleUnregisteredRoleName)
	if err != nil {
		v.Failf("fail", "%v", err)
	}
}
