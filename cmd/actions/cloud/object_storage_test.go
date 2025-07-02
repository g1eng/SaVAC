package cloud_actions

import (
	"context"
	"os"

	"github.com/g1eng/savac/pkg/cloud/sacloud"
	"github.com/g1eng/savac/pkg/core"
)

func (s *CloudActionSuite) TestScenario_O12_ListObjects() {
	if !isTestAcc {
		s.T().SkipNow()
	}
	testTable := []string{
		"s3://" + sampleObjectStorageBucketName,
		"s3://" + sampleObjectStorageBucketName + "/sample/path",
	}
	for _, p := range testTable {
		err = testAction(s.Generator.GenerateObjectStorageListActionMeta, p)
		if err != nil {
			s.Fail(err.Error())
		}
	}
}

func (s *CloudActionSuite) TestScenario_O12_ListSites_ListBuckets_Perm_PermKey_Create_Delete() {
	if !isTestAcc {
		s.T().SkipNow()
	}
	keyId := ""
	if isTestAcc {
		key, _, err := sacloudCli.ObjectStorageAPI.DefaultApi.CreateAccountKey(context.Background()).Execute()
		if err != nil {
			s.Fail(err.Error())
		} else if key == nil {
			s.Fail("no key created")
		}
		keyId = *key.Data.Id
		os.Setenv(core.OBJECT_STORAGE_API_KEY_ENV_NAME, keyId)
		os.Setenv(core.OBJECT_STORAGE_API_SECRET_ENV_NAME, *key.Data.Secret)
	}

	{
		cli, err := sacloud.NewCloudApiClient()
		if err != nil {
			s.Fail(err.Error())
		}
		newCommandGen := CloudActionGenerator{ApiClient: cli, Debug: true}
		err = testAction(newCommandGen.GenerateObjectStorageListSiteAction)
		if err != nil {
			s.Fail(err.Error())
		}
		err = testAction(newCommandGen.GenerateObjectStorageListBucketAction)
		if err != nil {
			s.Fail(err.Error())
		}
		dummyPerm := "test-perm-15"
		err = testAction(newCommandGen.GenerateCreatePermissionAction, "--rw", sampleObjectStorageBucketName, dummyPerm)
		if err != nil {
			s.Fail(err.Error())
		}
		err = testAction(newCommandGen.GenerateUpdatePermissionAction, "--wo", sampleObjectStorageBucketName, dummyPerm)
		if err != nil {
			s.Fail(err.Error())
		}
		err = testAction(newCommandGen.GenerateListPermissionAction, "--wo", sampleObjectStorageBucketName, dummyPerm)
		if err != nil {
			s.Fail(err.Error())
		}
		err = testAction(newCommandGen.GenerateCreatePermissionKeyAction, dummyPerm)
		if err != nil {
			s.Fail(err.Error())
		}
		err = testAction(newCommandGen.GenerateCreatePermissionKeyAction, dummyPerm)
		if err == nil {
			s.Fail("should fail for two or more keys")
		}
		err = testAction(newCommandGen.GenerateListPermissionKeyAction, dummyPerm)
		if err != nil {
			s.Fail(err.Error())
		}
		err = testAction(newCommandGen.GenerateDeletePermissionKeyAction, dummyPerm)
		if err != nil {
			s.Fail(err.Error())
		}
		err = testAction(newCommandGen.GenerateDeletePermissionKeyAction, dummyPerm)
		if err == nil {
			s.Fail("should fail for no key")
		}
		err = testAction(newCommandGen.GenerateDeletePermissionAction, dummyPerm)
		if err != nil {
			s.Fail(err.Error())
		}
		err = testAction(newCommandGen.GenerateDeletePermissionAction, dummyPerm)
		if err == nil {
			s.Fail("should fail for deleted permission id")
		}
	}

	if keyId != "" {
		_, err = sacloudCli.ObjectStorageAPI.DefaultApi.DeleteAccountKey(context.Background(), keyId).Execute()
		if err != nil {
			s.Fail(err.Error())
		}
	}
}

func (s *CloudActionSuite) TestScenario_O12_AccountKeyCreate_Delete() {
	if !isTestAcc {
		s.T().SkipNow()
	}
	err = testAction(s.Generator.GenerateCreateAccountKeyAction)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(s.Generator.GenerateDeleteAccountKeyAction)
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *CloudActionSuite) TestScenario_O12_Put_Check_Get_Objects() {
	if !isTestAcc {
		s.T().SkipNow()
	}
	testParentDir := "/tmp/savactest"
	testFilePath := testParentDir + "/savactestfile.txt"
	testDirPath := testParentDir + "/this/is"
	testFileDownloadPath := testParentDir + "/ok"
	testDirDownloadPath := testParentDir + "/ok-dir"
	err = os.MkdirAll(testDirPath, 0o0700)
	if err != nil {
		s.Fail(err.Error())
	}
	err := os.WriteFile(testFilePath, []byte("ok"), 0o0600)
	if err != nil {
		s.Fail(err.Error())
	}
	err = os.WriteFile(testDirPath+"/ok", []byte("ok"), 0o0600)
	if err != nil {
		s.Fail(err.Error())
	}
	testPutArguments := [][]string{
		{testFilePath, "s3://" + sampleObjectStorageBucketName + "/savac/test/sample/path.md"},
		{"--recursive", testDirPath, "s3://" + sampleObjectStorageBucketName + "/savac/test/sample/path"},
	}
	testCheckArguments := []string{"s3://" + sampleObjectStorageBucketName + "/savac/test/sample/path.md"}

	testGetArguments := [][]string{
		{"s3://" + sampleObjectStorageBucketName + "/savac/test/sample/path.md", testFileDownloadPath},
		{"--recursive", "s3://" + sampleObjectStorageBucketName + "/savac/test/sample/path", testDirDownloadPath},
	}
	testRemoveArguments := [][]string{
		{"s3://" + sampleObjectStorageBucketName + "/savac/test/sample/path.md"},
		{"--recursive", "s3://" + sampleObjectStorageBucketName + "/savac/test/sample/path"},
	}
	for _, p := range testPutArguments {
		err = testAction(s.Generator.GeneratePutAction, p...)
		if err != nil {
			s.Fail(err.Error())
		}
	}
	err = testAction(s.Generator.GenerateCheckAction, testCheckArguments...)
	if err != nil {
		s.Fail(err.Error())
	}
	for _, p := range testGetArguments {
		err = testAction(s.Generator.GenerateGetAction, p...)
		if err != nil {
			s.Fail(err.Error())
		}
		_, err := os.Stat(p[len(p)-1])
		if err != nil {
			s.Fail(err.Error())
		}
	}
	for _, p := range testRemoveArguments {
		err = testAction(s.Generator.GenerateRmAction, p...)
		if err != nil {
			s.Fail(err.Error())
		}
	}
	os.RemoveAll(testParentDir)
}
