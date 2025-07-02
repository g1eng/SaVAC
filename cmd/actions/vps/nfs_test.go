package vps_actions

import (
	"time"

	"github.com/g1eng/savac/testutil/test_parameters"
)

func (v *VpsActionSuite) TestGenerateNFSListAction() {
	err := testAction(actionGenerator.GenerateNfsListAction)
	if err != nil {
		v.Failf("fail", "should list nfs: %v", err)
	}
}

func (v *VpsActionSuite) TestGenerateNFSInfoAction() {
	err := testAction(actionGenerator.GenerateNfsInfoAction, test_parameters.SampleRegisteredNfsIdText)
	if err != nil {
		v.Failf("fail", "should show information of nfs: %v", err)
	}
	err = testAction(actionGenerator.GenerateNfsInfoAction, test_parameters.SampleRegisteredNfsIdText)
	if err != nil {
		v.Failf("fail", "should show information of nfs: %v", err)
	}
}

func (v *VpsActionSuite) TestGenerateNFSInfo_With_Nfs_Name_Action() {
	err := testAction(actionGenerator.GenerateNfsInfoAction, test_parameters.SampleRegisteredNfsName)
	if err != nil {
		v.Failf("fail", "should show information of nfs: %v", err)
	}
}

func (v *VpsActionSuite) TestGenerateNFSInterfaceAction() {
	err := testAction(actionGenerator.GenerateNfsInterfaceAction, test_parameters.SampleRegisteredNfsIdText)
	if err != nil {
		v.Failf("fail", "should list nfs interfaces: %v", err)
	}
}

func (v *VpsActionSuite) TestGenerateNFSConnectAction() {
	err := testAction(actionGenerator.GenerateNfsConnectAction, test_parameters.SampleRegisteredNfsIdText, test_parameters.SampleRegisteredNfsSwitchIdText)
	if err != nil {
		v.Failf("fail", "should connect nfs %s to the switch %s: %v", test_parameters.SampleRegisteredNfsIdText, test_parameters.SampleRegisteredNfsSwitchIdText, err)
	}
	err = testAction(actionGenerator.GenerateNfsConnectAction, "--disconnect", test_parameters.SampleRegisteredNfsIdText)
	if err != nil {
		v.Failf("fail", "should disconnect nfs %s from the switch %s: %v", test_parameters.SampleRegisteredNfsIdText, test_parameters.SampleRegisteredNfsSwitchIdText, err)
	}
}

func (v *VpsActionSuite) TestGenerateNFSConnect_With_Nfs_Name_Action() {
	err := testAction(actionGenerator.GenerateNfsConnectAction, test_parameters.SampleRegisteredNfsName, test_parameters.SampleRegisteredNfsSwitchIdText)
	if err != nil {
		v.Failf("fail", "should connect nfs %s to the switch %s: %v", test_parameters.SampleRegisteredNfsName, test_parameters.SampleRegisteredNfsSwitchIdText, err)
	}
	err = testAction(actionGenerator.GenerateNfsConnectAction, "--disconnect", test_parameters.SampleRegisteredNfsName)
	if err != nil {
		v.Failf("fail", "should disconnect nfs %s from the network: %v", test_parameters.SampleRegisteredNfsName, err)
	}
}

// This test assumes that the NFS is power on.
func (v *VpsActionSuite) TestGenerateNFSStartStopAction() {
	err := testAction(actionGenerator.GenerateNfsStartAction, test_parameters.SampleRegisteredNfsIdText)
	if err != nil {
		v.Failf("fail", "should power-on the nfs %s: %v", test_parameters.SampleRegisteredNfsIdText, err)
	}
	for _, args := range [][]string{
		{"--search", test_parameters.SampleRegisteredNfsName},
		{"--regex", test_parameters.SampleRegisteredNfsName + "o?"},
		{test_parameters.SampleRegisteredNfsName},
		{test_parameters.SampleRegisteredNfsIdText},
	} {
		err = testAction(actionGenerator.GenerateNfsStartAction, args...)
		if err != nil {
			v.Failf("fail", "should power-on the nfs by name %s: %v", args, err)
		}
		if v.isTestAcc {
			time.Sleep(5 * time.Second)
		}
		err = testAction(actionGenerator.GenerateNfsRebootAction, args...)
		if err != nil {
			v.Failf("fail", "should reboot the nfs %s: %v", args, err)
		}
		if v.isTestAcc {
			time.Sleep(15 * time.Second)
		}
		err = testAction(actionGenerator.GenerateNfsStopAction, args...)
		if err != nil {
			v.Failf("fail", "should power-off the nfs %s: %v", args, err)
		}
		if v.isTestAcc {
			time.Sleep(3 * time.Second)
		}
	}
}
