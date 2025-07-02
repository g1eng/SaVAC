package cloud_actions

import (
	"fmt"
	"time"

	"github.com/g1eng/savac/cmd/helper"
)

func (s *CloudActionSuite) TestScenario_WebAccel_CreateSubdomainSite_List_Enable_Disable_Delete() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	waSiteName := "neko2"
	err = testAction(s.Generator.GenerateWebAccelSiteCreateAction,
		"--origin",
		"docs.usacloud.jp",
		"--origin-protocol",
		"https",
		waSiteName,
	)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(s.Generator.GenerateWebAccelSiteListAction)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(s.Generator.GenerateWebAccelSiteUpdateStatusAction(true), waSiteName)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(s.Generator.GenerateWebAccelSiteUpdateStatusAction(false), waSiteName)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(s.Generator.GenerateWebAccelDeleteAction, waSiteName)
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *CloudActionSuite) TestScenario_WebAccel_CreateOwnDomainSite_Read_Update_Delete() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	waSiteName := "neko3"
	err = testAction(
		s.Generator.GenerateWebAccelSiteCreateAction,
		"--domain-type",
		"own_domain",
		"--domain",
		"nekowaf-tmp.savac.ns-testing.s0csec1.org",
		"--cors",
		"https://nekoneko.s0csec1.org",
		"--origin-type",
		"web",
		"--origin",
		"docs.usacloud.jp",
		"--host-header",
		"docs.usacloud.jp",
		"--origin-protocol",
		"https",
		"--request-protocol",
		"http+https",
		"--vary",
		"--accept-encoding",
		"brotli",
		"--default-cache-ttl",
		"3600",
		waSiteName,
	)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(s.Generator.GenerateWebAccelSiteReadAction, waSiteName)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(s.Generator.GenerateWebAccelSiteUpdateAction,
		"--cors",
		"https://nokenoke.com.com.com.s0csec1.org",
		"--origin-type",
		"bucket",
		"--endpoint", "s3.isk01.sakurastorage.jp",
		"--region", "jp-north-1",
		"--bucket", sampleObjectStorageBucketName,
		"--access-key", sampleObjectStorageAccessKey,
		"--access-secret", sampleObjectStorageAccessSecret,
		"--docindex",
		waSiteName)
	if err != nil {
		s.Fail(err.Error())
	}

	err = testAction(s.Generator.GenerateWebAccelDeleteAction, waSiteName)
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *CloudActionSuite) TestScenario_WebAccel_AutoRenew() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	err = testAction(s.Generator.GenerateWebAccelCertificateAutoRenewalAction, sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(s.Generator.GenerateWebAccelCertificateAutoRenewalAction, "--disable", sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *CloudActionSuite) TestScenario_WebAccel_CreateOnetimeSec_GenerateOnetimeURL_DeleteOneTimeSec() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	//set
	sec := helper.GenRandomString(16)
	err = testAction(s.Generator.GenerateWebAccelSiteWideOnetimeSecretAction, sampleWebAcceleratorSiteId, sec)
	if err != nil {
		s.Fail(err.Error())
	}
	//get (without arguments)
	err = testAction(s.Generator.GenerateWebAccelSiteWideOnetimeSecretAction, sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
	//generate one-time url with site id with default expiration time
	err = testAction(s.Generator.GenerateWebAccelOneTimeUrlAction, "--path", sampleWebAcceleratorURL, sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
	//generate one-time url with site url with default expiration time
	err = testAction(s.Generator.GenerateWebAccelOneTimeUrlAction, sampleWebAcceleratorURL)
	if err != nil {
		s.Fail(err.Error())
	}

	timeSpecTable := []string{
		"1mon",
		"1day",
		"1min",
		"1sec",
		"1hr",
		fmt.Sprintf("%d", time.Now().Add(time.Second*180).Unix()), //Unix time
	}
	//generate one-time url with specific expiration time
	for _, timeSpec := range timeSpecTable {
		err = testAction(s.Generator.GenerateWebAccelOneTimeUrlAction, "--expired", timeSpec, sampleWebAcceleratorURL)
		if err != nil {
			s.Fail(err.Error())
		}
	}

	//invalid
	err = testAction(s.Generator.GenerateWebAccelOneTimeUrlAction, "--expired", "neko", sampleWebAcceleratorURL)
	if err == nil {
		s.Fail("should have failed for invalid time spec: neko")
	}

	//set random secret
	err = testAction(s.Generator.GenerateWebAccelSiteWideOnetimeSecretAction, "--random", sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}

	//purge all secret and disable the site-wide one-time url
	err = testAction(s.Generator.GenerateWebAccelSiteWideOnetimeSecretAction, "--purge", sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *CloudActionSuite) TestScenario_WebAccel_OriginGuard() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	//create first
	err = testAction(s.Generator.GenerateWebAccelCreateOriginGuardTokenAction, sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
	//overwrite
	err = testAction(s.Generator.GenerateWebAccelCreateOriginGuardTokenAction, sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
	//next
	err = testAction(s.Generator.GenerateWebAccelCreateOriginGuardTokenAction, "--next", sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
	//apply next
	err = testAction(s.Generator.GenerateWebAccelCreateOriginGuardTokenAction, sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
	//next
	err = testAction(s.Generator.GenerateWebAccelCreateOriginGuardTokenAction, "--next", sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
	//cancel next
	err = testAction(s.Generator.GenerateWebAccelDeleteOriginGuardTokenAction, "--next", sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
	//delete all
	err = testAction(s.Generator.GenerateWebAccelDeleteOriginGuardTokenAction, sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *CloudActionSuite) TestWebAccel_PurgeCache() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	err = testAction(s.Generator.GenerateWebAccelPurgeCacheAction, sampleWebAcceleratorURL)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(s.Generator.GenerateWebAccelPurgeCacheAction, "all")
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *CloudActionSuite) TestScenario_WebAccel_Acl_Upsert_Read_Flush() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	err = testAction(s.Generator.GenerateWebAccelAclApplyAction, "--allow", "192.0.0.0/30", sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(s.Generator.GenerateWebAccelAclApplyAction, "--deny", "all", sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(s.Generator.GenerateWebAccelACLReadAction, sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
	err = testAction(s.Generator.GenerateWebAccelAclFlushAction, sampleWebAcceleratorSiteId)
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *CloudActionSuite) TestWebAccel_ApplyLogUpload() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	err = testAction(
		s.Generator.GenerateWebAccelAccessLogApplyAction,
		"--endpoint", "https://s3.isk01.sakurastorage.jp",
		"--region", "jp-north-1",
		"--bucket", sampleObjectStorageBucketName,
		"--access-key", sampleObjectStorageAccessKey,
		"--access-secret", sampleObjectStorageAccessSecret,
		sampleWebAcceleratorSiteId,
	)
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *CloudActionSuite) TestWebAccel_DeleteLogUpload() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	err = testAction(
		s.Generator.GenerateWebAccelAccessLogDeleteAction,
		sampleWebAcceleratorSiteId,
	)
	if err != nil {
		s.Fail(err.Error())
	}
}

func (s *CloudActionSuite) TestWebAccel_UsageRead() {
	if !isTestAcc {
		s.T().Skip("skip for non-acceptance test")
	}
	// for the month
	err = testAction(
		s.Generator.GenerateWebAccelUsageReadAction,
		"--month",
		fmt.Sprintf("%d", time.Now().Month()),
	)
	if err != nil {
		s.Fail(err.Error())
	}
	// for the month of the past year
	err = testAction(
		s.Generator.GenerateWebAccelUsageReadAction,
		"--year",
		fmt.Sprintf("%d", time.Now().Year()-1),
		"--month",
		fmt.Sprintf("%d", time.Now().Month()),
	)
	if err != nil {
		s.Fail(err.Error())
	}
	// for all past (and current) month of the year
	err = testAction(
		s.Generator.GenerateWebAccelUsageReadAction,
	)
	if err != nil {
		s.Fail(err.Error())
	}
}
