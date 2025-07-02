package cloud_actions

import (
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

var rId types.ID

func (s *CloudActionSuite) TestScenario_ContainerRegistry_Create_List_UserCreate_UserList_UserDelete_Delete() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	//CREATE
	err = testAction(s.Generator.GenerateContainerRegistryCreateAction, sampleUnregisteredContainerRegistryId)
	if err != nil {
		s.Fail(err.Error())
	}
	op := iaas.NewContainerRegistryOp(*s.Generator.ApiClient.Caller)
	rId, err = getRegistryIdByString(op, sampleUnregisteredContainerRegistryId)
	time.Sleep(10 * time.Second)

	//LIST
	err = testAction(s.Generator.GenerateContainerRegistryListAction, sampleUnregisteredContainerRegistryId)
	if err != nil {
		s.Fail(err.Error())
	}

	//USER CREATE
	err = testAction(
		s.Generator.GenerateContainerRegistryUserCreateAction,
		"--user", sampleUnregisteredContainerRegistryUserName,
		"--password", sampleUnregisteredContainerRegistryUserPassword,
		"--permission", "readonly",
		sampleUnregisteredContainerRegistryId,
	)
	if err != nil {
		s.Fail(err.Error())
	}
	time.Sleep(3 * time.Second)
	//USER LIST
	err = testAction(
		s.Generator.GenerateContainerRegistryUserListAction,
		sampleUnregisteredContainerRegistryId,
	)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(
		s.Generator.GenerateContainerRegistryUserDeleteAction,
		"--user",
		sampleUnregisteredContainerRegistryUserName,
		sampleUnregisteredContainerRegistryId,
	)
	if err != nil {
		s.Fail(err.Error())
	}
	time.Sleep(3 * time.Second)

	//USER DELETE
	err = testAction(s.Generator.GenerateContainerRegistryDeleteAction, rId.String())
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *CloudActionSuite) TestGenerateContainerRegistry_Fault_Create_Delete_Action() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	err = testAction(s.Generator.GenerateContainerRegistryCreateAction, sampleUnregisteredContainerRegistryId, "--permission", "readonly")
	if err == nil {
		s.Fail("should have failed without name flags")
	}
	err = testAction(s.Generator.GenerateContainerRegistryCreateAction, sampleUnregisteredContainerRegistryId, "--name", sampleUnregisteredContainerRegistryId, sampleUnregisteredContainerRegistryId)
	if err == nil {
		s.Fail("should have failed with a name flag and an argument")
	}
	err = testAction(s.Generator.GenerateContainerRegistryDeleteAction)
	if err == nil {
		s.Fail("should have failed without arguments")
	}
}
