package sacloud_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/g1eng/savac/pkg/cloud/sacloud"

	"github.com/g1eng/savac/pkg/cloud/model/object_storage"
	"github.com/g1eng/savac/pkg/core"
)

//func (s *ObjectStorageSuite) TestCloudAPIClient_CreateObjectStorageSiteAccount() {
//	err := s.Client.DeleteObjectStorageSiteAccount()
//	if err != nil {
//		s.Fail(err.Error())
//	}
//	if !isCI {
//		s.T().SkipNow()
//	}
//	_, err = s.Client.CreateObjectStorageSiteAccount()
//	if err != nil {
//		s.Fail(err.Error())
//	}
//}

//func (s *ObjectStorageSuite) TestCloudAPIClient_CreateBucket() {
//	if !isCI {
//		s.T().Skip()
//	}
//	_, err := s.Client.CreateBucket(newBucketName)
//	if err != nil {
//		s.Fail(err.Error())
//	}
//
//	//NOTE: fake server seems not to support DeleteBucket handler at now
//
//	//err = s.Client.DeleteBucket(newBucketName)
//	//if err != nil {
//	//	s.Fail(err.Error())
//	//}
//}

func (s *ObjectStorageSuite) TestCloudAPIClient_ListSites() {
	_, err = s.Client.ListSites()
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *ObjectStorageSuite) TestScenario_CloudAPIClient_CreateAccountKey_Create_List_Get_UpdatePerm_Create_List_DeletePermKey_DeletePerm_DeleteAccountKey() {
	res, err := s.Client.CreateAccountKey()
	if err != nil {
		s.Fail(err.Error())
	}
	SampleAccountKeyId = res.Data.Id
	if SampleAccountKeyId == nil {
		s.Fail("Account key not found")
	}
	SampleAccountSecret = res.Data.Secret
	time.Sleep(1 * time.Second)
	_, err = s.Client.ListAccountKeys()
	if err != nil {
		s.Fail(err.Error())
	}
	if isCI {
		os.Setenv(core.OBJECT_STORAGE_API_KEY_ENV_NAME, *SampleAccountKeyId)
		os.Setenv(core.OBJECT_STORAGE_API_SECRET_ENV_NAME, *SampleAccountSecret)
	}
	if err != nil {
		s.Fail(err.Error())
	}

	// List Buckets
	_, err = s.Client.ListBuckets()
	if err != nil {
		s.Fail(err.Error())
	}

	f := false
	resPerm, err := s.Client.CreatePermission(samplePermissionName, []object_storage.BucketControlsInner{
		{
			BucketName: &sampleBucketName,
			CanRead:    &f,
			CanWrite:   &f,
		},
	})
	if err != nil {
		s.Fail(err.Error())
	}
	samplePermissionId = fmt.Sprintf("%d", int(*resPerm.Data.Id))
	time.Sleep(1 * time.Second)
	listRes, err := s.Client.ListPermissions()
	if err != nil {
		s.Fail(err.Error())
	}
	if len(listRes) == 0 {
		s.Fail("no permissions found")
	}
	p, err := s.Client.GetPermissions(readonlyPermissionId)
	if err != nil {
		s.Fail(err.Error())
	}
	if fmt.Sprintf("%d", *p.Data.Id) != readonlyPermissionId {
		s.Failf("fail", "invlaid permission id fetched: %d", *p.Data.Id)
	}

	tr := true
	permControls := []object_storage.BucketControlsInner{
		{
			BucketName: &sampleBucketName,
			CanRead:    &tr,
			CanWrite:   &tr,
		},
	}
	_, err = s.Client.UpdatePermission(readonlyPermissionId, samplePermissionName, permControls)
	if err != nil {
		s.Fail(err.Error())
	}
	resPermkey, err := s.Client.CreatePermissionKey(readonlyPermissionId)
	if err != nil {
		s.Fail(err.Error())
	}
	if isCI == false {
		samplePermissionKeyId = *resPermkey.Data.Id
	}
	time.Sleep(1 * time.Second)
	resListPermKey, err := s.Client.ListPermissionKeys(readonlyPermissionId)
	if err != nil {
		s.Fail(err.Error())
	}
	if len(resListPermKey) == 0 {
		s.Fail("no permission keys found")
	}
	err = s.Client.DeletePermissionKey(readonlyPermissionId, samplePermissionKeyId)
	if err != nil {
		s.Fail(err.Error())
	}
	if isCI {
		samplePermissionId = readonlyPermissionId
	}
	err = s.Client.DeletePermission(samplePermissionId)
	if err != nil {
		s.Fail(err.Error())
	}
	if SampleAccountKeyId == nil {
		s.Fail("Account key not found")
	}
	if isCI {
		SampleAccountKeyId = &defaultAccountKeyId
	}
	err = s.Client.DeleteAccountKey(*SampleAccountKeyId)
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *ObjectStorageSuite) TestCloudAPIClient_For_FaultResponses() {
	n, m := "neko", "meko"
	tr := true
	go faultServer.ListenAndServe() // nolint
	time.Sleep(3 * time.Second)

	_, err := stubCli.CreateObjectStorageSiteAccount()
	if err == nil {
		s.Fail("should have failed")
	}
	_, err = stubCli.ListSites()
	if err == nil {
		s.Fail("should have failed")
	}
	err = stubCli.DeleteObjectStorageSiteAccount()
	if err == nil {
		s.Fail("should have failed")
	}
	_, err = stubCli.CreateAccountKey()
	if err == nil {
		s.Fail("should have failed")
	}
	_, err = stubCli.ListAccountKeys()
	if err == nil {
		s.Fail("should have failed")
	}
	err = stubCli.DeleteAccountKey(n)
	if err == nil {
		s.Fail("should have failed")
	}
	_, err = stubCli.CreateBucket(n)
	if err == nil {
		s.Fail("should have failed")
	}
	_, err = stubCli.ListBuckets()
	if err == nil {
		s.Fail("should have failed")
	}
	err = stubCli.DeleteBucket(n)
	if err == nil {
		s.Fail("should have failed")
	}
	_, err = stubCli.CreatePermission(n, []object_storage.BucketControlsInner{
		{
			BucketName: &n,
			CanRead:    &tr,
			CanWrite:   &tr,
		},
	})
	if err == nil {
		s.Fail("should have failed")
	}
	_, err = stubCli.ListPermissions()
	if err == nil {
		s.Fail("should have failed")
	}
	err = stubCli.DeletePermission(n)
	if err == nil {
		s.Fail("should have failed")
	}
	_, err = stubCli.CreatePermissionKey(n)
	if err == nil {
		s.Fail("should have failed")
	}
	_, err = stubCli.UpdatePermission(n, n, []object_storage.BucketControlsInner{
		{
			BucketName: &m,
			CanRead:    &tr,
			CanWrite:   &tr,
		},
	})
	if err == nil {
		s.Fail("should have failed")
	}
	_, err = stubCli.ListPermissionKeys(n)
	if err == nil {
		s.Fail("should have failed")
	}
	err = stubCli.DeletePermissionKey(n, n)
	if err == nil {
		s.Fail("should have failed")
	}
}

func (s *ObjectStorageSuite) TestDecomposeS3Uri() {
	validUri := []string{
		"s3://this/is/ok",
		"s3://that/is/ok/",
		"s3://those/are/ok/desu",
		"s3://bucket-without-key-is-valid",
	}
	bucketName := []string{
		"this",
		"that",
		"those",
		"bucket-without-key-is-valid",
	}

	key := []string{
		"is/ok",
		"is/ok",
		"are/ok/desu",
		"",
	}

	invalidUri := []string{
		"/this/is/wrong",
		"s3:/this/is/wrong",
		"https://this/is/wrong",
	}
	for i := range validUri {
		b, k, err := sacloud.DecomposeS3Uri(validUri[i])
		if err != nil {
			s.Failf("error returned", "uri %s should be valid: %s", validUri[i], err.Error())
		}
		if b != bucketName[i] {
			s.Failf("invalid bucket name", "bucket name %s should be %s", b, bucketName[i])
		}
		if k != key[i] {
			s.Failf("invalid key name", "key %s should be %s", k, key[i])
		}
	}
	for _, v := range invalidUri {
		_, _, err := sacloud.DecomposeS3Uri(v)
		if err == nil {
			s.Failf("URI parsing bug", "uri %s should be invalid", v)
		}
	}
}

func (s *ObjectStorageSuite) TestScenario_CloudAPIClient_Put_Check_Get_RemoveObject_RemoveRecursive() {
	if isCI {
		os.Setenv(core.OBJECT_STORAGE_API_KEY_ENV_NAME, *SampleAccountKeyId)
		os.Setenv(core.OBJECT_STORAGE_API_SECRET_ENV_NAME, *SampleAccountSecret)
	}

	// On macOS, this line may result invalid remote object path to be put
	err = s.Client.PutObject(sampleLocalObjectPath, sampleRemoteObjectParentKey)
	if err != nil {
		s.Fail(err.Error())
	}

	err = s.Client.PutObject(sampleLocalObjectPath, sampleRemoteObjectPath)
	if err != nil {
		s.Fail(err.Error())
	}
	err = s.Client.PutObject(sampleLocalObjectPath, sampleRemoteDeepObjectPath)
	if err != nil {
		s.Fail(err.Error())
	}

	err = s.Client.CheckObjectExistence(sampleRemoteObjectPath)
	if err != nil {
		s.Fail(err.Error())
	}

	r, err := s.Client.GetObject(sampleRemoteObjectPath)
	if err != nil {
		s.Fail(err.Error())
	}
	b, _ := io.ReadAll(r)
	b2, _ := os.ReadFile(sampleLocalObjectPath)
	if !bytes.Equal(b, b2) {
		s.Failf("files are different", "expected %s but got %s", string(b), string(b2))
	}

	res, err := s.Client.ListObjects(sampleRemoteObjectParentKey)
	if err != nil {
		s.Fail(err.Error())
	}
	if len(res) == 0 {
		s.Fail("response length should not be 0")
	}

	err = s.Client.RemoveObject(sampleRemoteObjectPath)
	if err != nil {
		s.Fail(err.Error())
	}
	err = s.Client.RemoveObject(sampleRemoteDeepObjectPath)
	if err != nil {
		s.Fail(err.Error())
	}
	err = s.Client.RemoveObjectRecursively(sampleRemoteObjectParentKey)
	if err != nil {
		s.Fail(err.Error())
	}
}

//func (s *ObjectStorageSuite) TestCloudAPIClient_uploadObjectInternal() {
//	if isCI {
//		os.Setenv(core.OBJECT_STORAGE_API_KEY_ENV_NAME, *SampleAccountKeyId)
//		os.Setenv(core.OBJECT_STORAGE_API_SECRET_ENV_NAME, *SampleAccountSecret)
//	}
//
//	// On macOS, this line may result invalid remote object path to be put
//	//err = newCli.uploadObjectInternal(sampleLocalObjectPath, sampleRemoteObjectParentKey)
//	f, err := os.Open(sampleLocalObjectPath)
//	if err != nil {
//		s.Fail(err.Error())
//	}
//	_, key, err := sacloud.DecomposeS3Uri(sampleRemoteObjectPath)
//	if err != nil {
//		s.Fail(err.Error())
//	}
//
//	err = s.Client.uploadObjectInternal(f, sampleBucketName, key)
//	if err != nil {
//		s.Fail(err.Error())
//	}
//
//	_, key, err = sacloud.DecomposeS3Uri(sampleRemoteAnotherObjectPath)
//	if err != nil {
//		s.Fail(err.Error())
//	}
//	err = s.Client.uploadObjectInternal(f, sampleBucketName, key)
//	if err != nil {
//		s.Fail(err.Error())
//	}
//}
