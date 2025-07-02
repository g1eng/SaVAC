package cloud_actions

import (
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/g1eng/savac/pkg/core"

	"github.com/g1eng/savac/cmd/helper"
	"github.com/g1eng/savac/pkg/cloud/sacloud"
	"github.com/sacloud/webaccel-api-go"
	"github.com/urfave/cli/v3"
)

func buildSiteRequest(cmd *cli.Command, req *webaccel.UpdateSiteRequest) (*webaccel.UpdateSiteRequest, error) {
	switch cmd.String("request-protocol") {
	case "http+https":
		req.RequestProtocol = webaccel.RequestProtocolsHttpAndHttps
	case "https":
		req.RequestProtocol = webaccel.RequestProtocolsHttpsOnly
	case "https-redirect":
		req.RequestProtocol = webaccel.RequestProtocolsRedirectToHttps
	}
	if cmd.IsSet("default-cache-ttl") {
		i := cmd.Int("default-cache-ttl")
		req.DefaultCacheTTL = &i
	}
	corsOrigins := cmd.StringSlice("cors")
	if len(corsOrigins) != 0 {
		corsRules := []*webaccel.CORSRule{{AllowsAnyOrigin: false}}
		if strings.Join(corsOrigins, "") == "*" {
			corsRules = []*webaccel.CORSRule{{AllowsAnyOrigin: true}}
		} else {
			corsRules[0].AllowedOrigins = corsOrigins
		}
		req.CORSRules = &corsRules
	}

	if cmd.IsSet("vary") {
		if cmd.Bool("vary") {
			req.VarySupport = webaccel.VarySupportEnabled
		} else {
			req.VarySupport = webaccel.VarySupportDisabled
		}
	}
	if cmd.String("accept-encoding") == "brotli" || cmd.String("accept-encoding") == "bz+gzip" {
		req.NormalizeAE = webaccel.NormalizeAEBrGz
	} else if cmd.String("accept-encoding") == "gzip" {
		req.NormalizeAE = webaccel.NormalizeAEGz
	}

	switch cmd.String("origin-type") {
	case "bucket":
		req.OriginType = webaccel.OriginTypesObjectStorage
		if cmd.IsSet("bucket") {
			req.BucketName = cmd.String("bucket")
		}
		if cmd.IsSet("endpoint") {
			req.S3Endpoint = cmd.String("endpoint")
		}
		if cmd.IsSet("region") {
			req.S3Region = cmd.String("region")
		}
		if cmd.IsSet("access-key") {
			req.AccessKeyID = cmd.String("access-key")
		}
		if cmd.IsSet("access-secret") {
			req.SecretAccessKey = cmd.String("access-secret")
		}
		if cmd.IsSet("docindex") {
			if cmd.Bool("docindex") {
				req.DocIndex = webaccel.DocIndexEnabled
			} else {
				req.DocIndex = webaccel.DocIndexDisabled
			}
		}
		if cmd.IsSet("origin") || cmd.IsSet("host-header") {
			return nil, fmt.Errorf("invalid flag specified for bucket origin")
		}
	case "web":
		fallthrough
	default:
		req.OriginType = webaccel.OriginTypesWebServer
		if cmd.IsSet("origin") {
			req.Origin = cmd.String("origin")
		}
		if cmd.IsSet("origin-protocol") {
			req.OriginProtocol = cmd.String("origin-protocol")
		}
		if cmd.IsSet("host-header") {
			req.HostHeader = cmd.String("host-header")
		}
		if cmd.IsSet("bucket") || cmd.IsSet("docindex") {
			return nil, fmt.Errorf("invalid flag specified for web origin")
		}
	}
	return req, nil
}

func (g *CloudActionGenerator) normalizeSiteId(idCandidate string) (string, error) {
	_, err := strconv.Atoi(idCandidate)
	if err == nil {
		return idCandidate, nil
	}
	res, err := g.getWebAccelSiteWithName(idCandidate)
	if err != nil {
		return "", err
	}
	return res.ID, nil
}
func (g *CloudActionGenerator) getWebAccelSiteWithName(name string) (*webaccel.Site, error) {
	res, err := g.ApiClient.WebAccelAPI.List(context.Background())
	if err != nil {
		return nil, err
	}
	if res.Count == 0 {
		return nil, fmt.Errorf("no sites")
	}
	for _, r := range res.Sites {
		if r.Name == name {
			return r, nil
		} else if r.ASCIIDomain == name {
			return r, nil
		}
	}
	return nil, fmt.Errorf("site %s not found", name)
}
func (g *CloudActionGenerator) GenerateWebAccelSiteCreateAction(c context.Context, cmd *cli.Command) error {
	requestCoreParams, err := buildSiteRequest(cmd, &webaccel.UpdateSiteRequest{})
	if err != nil {
		return err
	}
	if cmd.String("origin-type") == "web" {
		if !cmd.IsSet("origin") {
			return fmt.Errorf("--origin flag is required for web origin type")
		}
	} else if cmd.String("origin-type") == "bucket" {
		if !cmd.IsSet("bucket") {
			return fmt.Errorf("--bucket flag is required for bucket origin type")
		} else if !cmd.IsSet("access-key") {
			return fmt.Errorf("--access-key flag is required for bucket origin type")
		} else if !cmd.IsSet("access-secret") {
			return fmt.Errorf("--access-secret flag is required for bucket origin type")
		} else if !cmd.IsSet("region") {
			return fmt.Errorf("--region flag is required for bucket origin type")
		} else if !cmd.IsSet("endpoint") {
			return fmt.Errorf("--endpoint flag is required for bucket origin type")
		}
	} else {
		return fmt.Errorf("unknown origin type: %s", cmd.String("origin-type"))
	}
	req := &webaccel.CreateSiteRequest{
		Name:            cmd.Args().First(),
		DomainType:      cmd.String("domain-type"),
		RequestProtocol: requestCoreParams.RequestProtocol,
		VarySupport:     requestCoreParams.VarySupport,
		DefaultCacheTTL: requestCoreParams.DefaultCacheTTL,
		OriginType:      requestCoreParams.OriginType,
		OriginProtocol:  requestCoreParams.OriginProtocol,
		Origin:          requestCoreParams.Origin,
		HostHeader:      requestCoreParams.HostHeader,
		S3Endpoint:      requestCoreParams.S3Endpoint,
		S3Region:        requestCoreParams.S3Region,
		BucketName:      requestCoreParams.BucketName,
		DocIndex:        requestCoreParams.DocIndex,
		AccessKeyID:     requestCoreParams.AccessKeyID,
		SecretAccessKey: requestCoreParams.SecretAccessKey,
	}
	if cmd.String("domain-type") == "own_domain" {
		req.Domain = cmd.String("domain")
	}

	resp, err := g.ApiClient.WebAccelAPI.Create(context.Background(), req)
	if err != nil {
		return err
	}
	return helper.PrintJson(resp)
}

func (g *CloudActionGenerator) GenerateWebAccelSiteListAction(c context.Context, cmd *cli.Command) error {
	resp, err := g.ApiClient.WebAccelAPI.List(context.Background())
	if err != nil {
		return err
	}
	return helper.PrintJson(resp)
}

func (g *CloudActionGenerator) GenerateWebAccelSiteReadAction(c context.Context, cmd *cli.Command) error {
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	resp, err := g.ApiClient.WebAccelAPI.Read(context.Background(), id)
	if err != nil {
		return err
	}
	return helper.PrintJson(resp)
}

func (g *CloudActionGenerator) GenerateWebAccelSiteUpdateAction(_ context.Context, cmd *cli.Command) error {
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	req, err := buildSiteRequest(cmd, &webaccel.UpdateSiteRequest{})
	if err != nil {
		return err
	}
	res, err := g.ApiClient.WebAccelAPI.Update(context.Background(), id, req)
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}

func (g *CloudActionGenerator) GenerateWebAccelSiteWideOnetimeSecretAction(c context.Context, cmd *cli.Command) error {
	var site *webaccel.Site
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	site, err = g.ApiClient.WebAccelAPI.Read(context.Background(), id)
	if err != nil {
		return err
	}
	if cmd.Bool("purge") {
		site.OnetimeURLSecrets = []string{}
	} else if cmd.NArg() > 1 {
		site.OnetimeURLSecrets = append(site.OnetimeURLSecrets, cmd.Args().Slice()[1])
	} else if cmd.Bool("random") {
		shiftLength := rand.Intn(5) // nolint
		site.OnetimeURLSecrets = append(site.OnetimeURLSecrets, helper.GenRandomString(12+shiftLength))
	} else {
		return helper.PrintJson(site.OnetimeURLSecrets)
	}
	if len(site.OnetimeURLSecrets) > 2 {
		site.OnetimeURLSecrets = site.OnetimeURLSecrets[1:]
	}
	res, err := g.ApiClient.WebAccelAPI.Update(context.Background(), site.ID, &webaccel.UpdateSiteRequest{
		OnetimeURLSecrets: &site.OnetimeURLSecrets,
	})
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}

func (g *CloudActionGenerator) GenerateWebAccelOneTimeUrlAction(c context.Context, cmd *cli.Command) (err error) {
	var (
		u    *url.URL
		site *webaccel.Site
		path string
	)
	desc := cmd.Args().First()
	if strings.Contains(desc, ".") {
		u, err = url.Parse(desc)
		if err != nil {
			return err
		}
		site, err = g.getWebAccelSiteWithName(u.Host)
		if err != nil {
			return err
		}
		site2, err := g.ApiClient.WebAccelAPI.Read(context.Background(), site.ID)
		if err != nil {
			return err
		}
		site.OnetimeURLSecrets = site2.OnetimeURLSecrets
		path = u.Path
	} else {
		u = &url.URL{}
		if !cmd.IsSet("path") {
			return fmt.Errorf("path must be set for the non-url argument")
		}
		_, err := strconv.Atoi(desc)
		if err != nil {
			site, err = g.getWebAccelSiteWithName(desc)
		} else {
			site, err = g.ApiClient.WebAccelAPI.Read(context.Background(), desc)
		}
		if err != nil {
			return err
		}
		path = cmd.String("path")
	}
	var expirationTime time.Time
	at := cmd.String("expired")
	if regexp.MustCompilePOSIX("^[[:digit:]]+(mon|day|hr|min|sec)$").MatchString(at) {
		if strings.Contains(at, "mon") {
			multi, _ := strconv.Atoi(strings.TrimRight(at, "mon"))
			expirationTime = time.Now().Add(time.Hour * 24 * 30 * time.Duration(multi))
		} else if strings.Contains(at, "day") {
			multi, _ := strconv.Atoi(strings.TrimRight(at, "day"))
			expirationTime = time.Now().Add(time.Hour * 24 * time.Duration(multi))
		} else if strings.Contains(at, "hr") {
			multi, _ := strconv.Atoi(strings.TrimRight(at, "hr"))
			expirationTime = time.Now().Add(time.Hour * time.Duration(multi))
		} else if strings.Contains(at, "min") {
			multi, _ := strconv.Atoi(strings.TrimRight(at, "min"))
			expirationTime = time.Now().Add(time.Minute * time.Duration(multi))
		} else if strings.Contains(at, "sec") {
			multi, _ := strconv.Atoi(strings.TrimRight(at, "sec"))
			expirationTime = time.Now().Add(time.Second * time.Duration(multi))
		}
	} else if unix, err := strconv.Atoi(at); err != nil {
		return err
	} else {
		expirationTime = time.Unix(int64(unix), 0)
	}
	if len(site.OnetimeURLSecrets) == 0 {
		return fmt.Errorf("no secret found")
	}
	path = sacloud.GenerateOnetimePath(site.OnetimeURLSecrets[len(site.OnetimeURLSecrets)-1], path, expirationTime)
	// fmt.Printf("original %s\n", desc)
	fmt.Printf("https://%s%s\n", site.ASCIIDomain, path)

	return nil
}

func (g *CloudActionGenerator) GenerateWebAccelACLReadAction(c context.Context, cmd *cli.Command) error {
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	res, err := g.ApiClient.WebAccelAPI.ReadACL(context.Background(), id)
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}

func (g *CloudActionGenerator) GenerateWebAccelAclApplyAction(_ context.Context, cmd *cli.Command) error {
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	rules := ""
	if !cmd.IsSet("allow") && !cmd.IsSet("deny") {
		return fmt.Errorf("no rule specified")
	}
	for _, r := range cmd.StringSlice("allow") {
		rules += fmt.Sprintf("allow %s\n", r)
	}
	if !cmd.IsSet("deny") {
		rules += "deny all\n"
	}
	for _, r := range cmd.StringSlice("deny") {
		rules += fmt.Sprintf("deny %s\n", r)
	}
	res, err := g.ApiClient.WebAccelAPI.UpsertACL(context.Background(), id, rules)
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}

func (g *CloudActionGenerator) GenerateWebAccelAclFlushAction(c context.Context, cmd *cli.Command) error {
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	return g.ApiClient.WebAccelAPI.DeleteACL(context.Background(), id)
}

func (g *CloudActionGenerator) GenerateWebAccelAccessLogApplyAction(_ context.Context, cmd *cli.Command) error {
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	region := cmd.String("region")
	endpoint := cmd.String("endpoint")
	bucketName := cmd.String("bucket")
	accessKey := cmd.String("access-key")
	accessSecret := cmd.String("access-secret")
	param := webaccel.LogUploadConfig{
		Bucket:          bucketName,
		Endpoint:        endpoint,
		Region:          region,
		AccessKeyID:     accessKey,
		SecretAccessKey: accessSecret,
		Status:          "disabled",
	}
	res, err := g.ApiClient.WebAccelAPI.ApplyLogUploadConfig(context.Background(), id, &param)
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}

func (g *CloudActionGenerator) GenerateWebAccelAccessLogReadAction(_ context.Context, cmd *cli.Command) error {
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	res, err := g.ApiClient.WebAccelAPI.ReadLogUploadConfig(context.Background(), id)
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}

func (g *CloudActionGenerator) GenerateWebAccelAccessLogUpdateStatusAction(toEnable bool) func(context.Context, *cli.Command) error {
	return func(c context.Context, cmd *cli.Command) error {
		id, err := g.normalizeSiteId(cmd.Args().First())
		if err != nil {
			return err
		}
		conf, err := g.ApiClient.WebAccelAPI.ReadLogUploadConfig(context.Background(), id)
		if err != nil {
			return err
		}
		if conf.Bucket == "" {
			return fmt.Errorf("no configuration")
		}
		statusString := "disabled"
		if toEnable {
			statusString = "enabled"
		}
		param := webaccel.LogUploadConfig{
			Bucket:          conf.Bucket,
			Endpoint:        conf.Endpoint,
			Region:          conf.Region,
			AccessKeyID:     conf.AccessKeyID,
			SecretAccessKey: conf.SecretAccessKey,
			Status:          statusString,
		}
		res, err := g.ApiClient.WebAccelAPI.ApplyLogUploadConfig(context.Background(), id, &param)
		if err != nil {
			return err
		}
		return helper.PrintJson(res)
	}
}

func (g *CloudActionGenerator) GenerateWebAccelAccessLogDeleteAction(c context.Context, cmd *cli.Command) error {
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	return g.ApiClient.WebAccelAPI.DeleteLogUploadConfig(context.Background(), id)
}

func (g *CloudActionGenerator) GenerateWebAccelPurgeCacheAction(c context.Context, cmd *cli.Command) error {
	if cmd.Bool("all") {
		for i, d := range cmd.Args().Slice() {
			if i != 0 {
				time.Sleep(time.Second)
			}
			err := g.ApiClient.WebAccelAPI.DeleteAllCache(context.Background(), &webaccel.DeleteAllCacheRequest{Domain: d})
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		u := cmd.Args().Slice()
		res, err := g.ApiClient.WebAccelAPI.DeleteCache(context.Background(), &webaccel.DeleteCacheRequest{URL: u})
		if err != nil {
			return err
		}
		return helper.PrintJson(res)
	}
}

func (g *CloudActionGenerator) GenerateWebAccelUsageReadAction(c context.Context, cmd *cli.Command) error {
	if !cmd.IsSet("year") {
		err := cmd.Set("year", fmt.Sprintf("%d", time.Now().Year()))
		if err != nil {
			return err
		}
	}
	if cmd.IsSet("month") {
		use, err := g.ApiClient.WebAccelAPI.MonthlyUsage(context.Background(), fmt.Sprintf("%d%02d", cmd.Int("year"), cmd.Int("month")))
		if err != nil {
			return err
		}
		return helper.PrintJson(use)
	} else {
		limit := 12
		if cmd.Int("year") == time.Now().Year() {
			limit = int(time.Now().Month())
		}

		for i := 1; i <= limit; i++ {
			if i != 1 {
				time.Sleep(time.Millisecond * 350)
			}
			use, err := g.ApiClient.WebAccelAPI.MonthlyUsage(context.Background(), fmt.Sprintf("%d%02d", cmd.Int("year"), i))
			if err != nil {
				return err
			}
			helper.PrintJson(use) //nolint
		}
		return nil
	}
}

func (g *CloudActionGenerator) GenerateWebAccelDeleteAction(c context.Context, cmd *cli.Command) error {
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	resp, err := g.ApiClient.WebAccelAPI.Delete(context.Background(), id)
	if err != nil {
		return err
	}
	return helper.PrintJson(resp)
}

func (g *CloudActionGenerator) GenerateWebAccelSiteUpdateStatusAction(toEnable bool) func(context.Context, *cli.Command) error {
	return func(c context.Context, cmd *cli.Command) error {
		id, err := g.normalizeSiteId(cmd.Args().First())
		if err != nil {
			return err
		}
		statusString := "enabled"
		if !toEnable {
			statusString = "disabled"
		}
		req := &webaccel.UpdateSiteStatusRequest{
			Status: statusString,
		}
		resp, err := g.ApiClient.WebAccelAPI.UpdateStatus(context.Background(), id, req)
		if err != nil {
			return err
		}
		return helper.PrintJson(resp)
	}
}

// FIXME: WIP
func (g *CloudActionGenerator) GenerateWebAccelReadOriginGuardTokenAction(context.Context, *cli.Command) error {
	return fmt.Errorf("WIP")
}

func (g *CloudActionGenerator) GenerateWebAccelCreateOriginGuardTokenAction(_ context.Context, cmd *cli.Command) error {
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	var res *webaccel.OriginGuardTokenResponse
	if cmd.Bool("next") {
		res, err = g.ApiClient.WebAccelAPI.CreateNextOriginGuardToken(context.Background(), id)
	} else {
		res, err = g.ApiClient.WebAccelAPI.CreateOriginGuardToken(context.Background(), id)
	}
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}

func (g *CloudActionGenerator) GenerateWebAccelDeleteOriginGuardTokenAction(ctx context.Context, cmd *cli.Command) error {
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	if cmd.Bool("next") {
		return g.ApiClient.WebAccelAPI.DeleteNextOriginGuardToken(context.Background(), id)
	} else {
		return g.ApiClient.WebAccelAPI.DeleteOriginGuardToken(context.Background(), id)
	}
}

// GenerateWebAccelCertificateImportAction generates certificate import action
// FIXME: test it automatically
func (g *CloudActionGenerator) GenerateWebAccelCertificateImportAction(_ context.Context, cmd *cli.Command) error {
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	cert, err := os.ReadFile(cmd.String("file"))
	if err != nil {
		return err
	}
	//cert = []byte(strings.TrimRight(string(cert), "\n"))
	req := &webaccel.CreateOrUpdateCertificateRequest{
		CertificateChain: string(cert),
	}
	if cmd.IsSet("key") {
		if cmd.String("key") != "" {
			key, err := os.ReadFile(cmd.String("key"))
			if err != nil {
				return err
			}
			req.Key = string(key)
		}
	}
	res, err := g.ApiClient.WebAccelAPI.CreateCertificate(context.Background(), id, req)
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}

// GenerateWebAccelCertificateReadAction generates action for certificate read command
// FIXME: test it automatically
func (g *CloudActionGenerator) GenerateWebAccelCertificateReadAction(_ context.Context, cmd *cli.Command) error {
	res, err := g.ApiClient.WebAccelAPI.ReadCertificate(context.Background(), cmd.Args().First())
	if err != nil {
		return err
	}
	if res.Current == nil {
		return fmt.Errorf("no certificates")
	}

	var cert interface{}
	switch i := cmd.Int("revision"); i {
	case 0:
		fallthrough
	case len(res.Old) + 1:
		cert = res.Current
	default:
		if i > len(res.Old)-1 {
			return fmt.Errorf("revision out of range")
		}
		cert = res.Old[i]
	}

	switch g.OutputType {
	case core.OutputTypeText:
		if curCert, ok := cert.(*webaccel.CurrentCertificate); ok {
			println(curCert.CertificateChain)
		} else if oldCert, ok := cert.(*webaccel.OldCertificate); ok {
			println(oldCert.CertificateChain)
		} else {
			panic("invalid condition: certificate type parse error")
		}
		return nil
	case core.OutputTypeYaml:
		return helper.PrintYaml(cert)
	case core.OutputTypeJson:
		fallthrough
	default:
		return helper.PrintYaml(cert)
	}
}

func (g *CloudActionGenerator) GenerateWebAccelRevisionsAction(_ context.Context, cmd *cli.Command) error {
	res, err := g.ApiClient.WebAccelAPI.ReadCertificate(context.Background(), cmd.Args().First())
	if err != nil {
		return err
	}
	switch g.OutputType {
	case core.OutputTypeYaml:
		return helper.PrintYaml(res)
	case core.OutputTypeText:
		return helper.PrintTableForWebaccelCertificateRevisions(res)
	case core.OutputTypeJson:
		fallthrough
	default:
		return helper.PrintJson(res)
	}
}

func (g *CloudActionGenerator) GenerateWebAccelCertificateAutoRenewalAction(ctx context.Context, cmd *cli.Command) error {
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	if !cmd.Bool("lets-encrypt") {
		return fmt.Errorf("auto renewal is only support for LetsEncrypt")
	}
	if cmd.Bool("disable") {
		return g.ApiClient.WebAccelAPI.DeleteAutoCertUpdate(context.Background(), id)
	}
	err = g.ApiClient.WebAccelAPI.CreateAutoCertUpdate(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}

func (g *CloudActionGenerator) GenerateWebAccelCertificateDeleteAction(ctx context.Context, cmd *cli.Command) error {
	id, err := g.normalizeSiteId(cmd.Args().First())
	if err != nil {
		return err
	}
	return g.ApiClient.WebAccelAPI.DeleteCertificate(ctx, id)
}
