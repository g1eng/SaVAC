package vps

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/g1eng/savac/testutil/test_parameters"
)

func (v *SavaClientSuite) TestSavaClient_GetAllNFS() {
	res, err := v.nfsCli.GetAllNFS()
	if err != nil {
		v.Fail("%v", err)
	}
	if len(res) == 0 {
		v.Fail("NFS server should be listed")
	}
}

func (v *SavaClientSuite) TestSavaClient_ListNFSInterfaces() {
	res, err := v.nfsCli.GetNFSInterface(test_parameters.SampleRegisteredNfsId)
	if err != nil {
		v.Fail("%v", err)
	}
	if res == nil {
		v.Fail("NFS server interface should be given")
	}
}

func (v *SavaClientSuite) TestSavaClient_GetNFSById() {
	n, err := v.nfsCli.GetNFSById(test_parameters.SampleRegisteredNfsId)
	if err != nil {
		v.Fail("%v", err)
	}
	if n.Id != test_parameters.SampleRegisteredNfsId {
		v.Fail("NFS id should be identical to %d", test_parameters.SampleRegisteredNfsId)
	}
}

func (v *SavaClientSuite) TestSavaClient_SetNFSDescription() {
	err := v.nfsCli.SetNFSDescription(test_parameters.SampleRegisteredNfsId, "ok")
	v.NoError(err)
}

func (v *SavaClientSuite) TestSavaClient_SetNFSName() {
	err := v.nfsCli.SetNFSName(test_parameters.SampleRegisteredNfsId, test_parameters.SampleRegisteredNfsName)
	v.NoError(err)
}

func (v *SavaClientSuite) TestSavaClient_GetNFSByName() {
	nfs, err := v.nfsCli.GetNFSByName(test_parameters.SampleRegisteredNfsName)
	if err != nil {
		v.Fail("%v", err)
	}
	if nfs.Name == test_parameters.SampleRegisteredNfsName {
		return
	}
	v.Fail("The NFS name %s should be contained in the result", test_parameters.SampleRegisteredNfsName)
}

func (v *SavaClientSuite) TestSavaClient_GetNFSByRegex() {
	reg := regexp.MustCompilePOSIX("^" + test_parameters.SampleRegisteredNfsName + "$")
	res, err := v.nfsCli.GetNFSListByRegex(reg)
	if err != nil {
		v.Fail("%v", err)
	}
	if len(res) == 0 {
		v.Fail("NFS server should exist")
	}
	if res[0].Name == test_parameters.SampleRegisteredNfsName {
		return
	}
	v.Fail("The NFS name %s should be contained in the result", test_parameters.SampleRegisteredNfsName)
}

func (v *SavaClientSuite) TestSavaClient_FaultNFSRequests() {
	if v.isTestAcc {
		time.Sleep(time.Second)
	}
	_, err := v.faultCli.GetAllNFS()
	if err == nil {
		v.Fail("should raise an exception for server fault")
	}
	_, err = v.faultCli.GetNFSInterface(test_parameters.SampleRegisteredNfsId)
	if err == nil {
		v.Fail("should raise an exception for server fault")
	}
	_, err = v.faultCli.GetNFSById(test_parameters.SampleRegisteredNfsId)
	if err == nil {
		v.Fail("should raise an exception for server fault")
	}
	_, err = v.faultCli.GetNFSByName(test_parameters.SampleRegisteredNfsName)
	if err == nil {
		v.Fail("should raise an exception for server fault")
	}
	err = v.faultCli.SetNFSName(test_parameters.SampleRegisteredNfsId, test_parameters.SampleRegisteredNfsName)
	if err == nil {
		v.Fail("should raise an exception for server fault")
	}
	err = v.faultCli.SetNFSDescription(test_parameters.SampleRegisteredNfsId, test_parameters.SampleRegisteredNfsName)
	if err == nil {
		v.Fail("should raise an exception for server fault")
	}
	err = v.faultCli.PutNFSConnection(test_parameters.SampleRegisteredNfsId, test_parameters.SampleRegisteredNfsSwitchId)
	if err == nil {
		v.Fail("should raise an exception for server fault")
	}
}

func (v *SavaClientSuite) TestScenario_SavaClient_GetNFSStorageInfo() {
	res, _, err := v.nfsCli.RawClient.NfsServerAPI.GetNfsServerPowerStatus(context.Background(), test_parameters.SampleRegisteredNfsId).Execute()
	v.NoError(err)
	if res.Status == "power_on" {
		_, err := v.nfsCli.RawClient.NfsServerAPI.PostNfsServerShutdown(context.Background(), test_parameters.SampleRegisteredNfsId).Execute()
		v.NoError(err)
		if v.isTestAcc {
			time.Sleep(time.Second * 10)
		}
	}
	//off
	err = v.nfsCli.PutNFSConnection(test_parameters.SampleRegisteredNfsId, test_parameters.SampleRegisteredNfsSwitchId)
	v.NoError(err)
	err = v.nfsCli.PutNFSConnection(test_parameters.SampleRegisteredNfsId, 0)
	v.NoErrorf(err, "failed to disconnect NFS server from a switch")

	_, err = v.nfsCli.RawClient.NfsServerAPI.PostNfsServerPowerOn(context.Background(), test_parameters.SampleRegisteredNfsId).Execute()
	v.NoError(err)
	if v.isTestAcc {
		time.Sleep(time.Second * 20)
	}

	//on
	info, err := v.nfsCli.GetNfsStorageInfo(test_parameters.SampleRegisteredNfsId)
	if err != nil {
		v.Fail("%v", err)
	}
	if info.GetCapacityKib() < info.GetUsageKib() {
		v.Fail("invalid capacity returned for the nfs: %d", test_parameters.SampleRegisteredNfsId)
	}
	fmt.Printf("use %d / cap %d (%d %s)\n", info.GetUsageKib(), info.GetCapacityKib(), info.GetUsagePercentage(), "%")
}

func (v *SavaClientSuite) TestScenario_SavaClient_Shutdown_Start_Reboot_ForciblyShutdown_NFS() {
	res, _, err := v.nfsCli.RawClient.NfsServerAPI.GetNfsServerPowerStatus(context.Background(), test_parameters.SampleRegisteredNfsId).Execute()
	if err != nil {
		v.Fail("failed to get Nfs power status: %v", err)
	}
	if res.Status == "power_off" {
		_, err := v.nfsCli.RawClient.NfsServerAPI.PostNfsServerPowerOn(context.Background(), test_parameters.SampleRegisteredNfsId).Execute()
		v.NoError(err)
		if v.isTestAcc {
			time.Sleep(time.Second * 10)
		}
	}
	err = v.nfsCli.ShutdownNfs(test_parameters.SampleRegisteredNfsId)
	if err != nil {
		v.Fail("%v", err)
	}
	if v.isTestAcc {
		time.Sleep(time.Second * 5)
	}

	//start
	err = v.nfsCli.StartNfs(test_parameters.SampleRegisteredNfsId)
	if err != nil {
		v.Fail("%v", err)
	}
	if v.isTestAcc {
		time.Sleep(time.Second * 20)
	}

	//force reboot
	err = v.nfsCli.RebootNfs(test_parameters.SampleRegisteredNfsId)
	if err != nil {
		v.Fail("%v", err)
	}
	if v.isTestAcc {
		time.Sleep(time.Second * 5)
	}

	//force shutdown
	err = v.nfsCli.ShutdownNfs(test_parameters.SampleRegisteredNfsId, true)
	if err != nil {
		v.Fail("%v", err)
	}
}
