package fake_cloud

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	objectStorageFakeApi "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/sacloud/object-storage-api-go/fake"
	"github.com/sacloud/object-storage-api-go/fake/server"
)

const (
	defaultBucketName       = "dummy-bucket-name-this-id"
	samplePermissionName    = "test-savac-perm"
	readonlyPermissionIdInt = 8917
	readonlyPermissionKeyId = "blank"
)

func NewObjectStorageFakeHandler(fakeEngine *fake.Engine) http.Handler {
	fakeServer := &server.Server{
		Engine: fakeEngine,
	}
	return fakeServer.Handler()
}

type MappedFakeHandler struct {
	Handler http.Handler
}

func (h *MappedFakeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	doubleSlash := regexp.MustCompile("//+")
	r.URL.Path, _ = strings.CutPrefix(r.URL.Path, "/cloud/zone/is1a/api/objectstorage/1.0")
	r.URL.Path = doubleSlash.ReplaceAllString(r.URL.Path, "/")
	r.RequestURI, _ = strings.CutPrefix(r.URL.Path, "/cloud/zone/is1a/api/objectstorage/1.0")
	r.RequestURI = doubleSlash.ReplaceAllString(r.URL.Path, "/")
	h.Handler.ServeHTTP(w, r)
}

func NewObjectStorageFakeServer(port int, engine *fake.Engine) *http.Server {
	sv := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: &MappedFakeHandler{
			Handler: NewObjectStorageFakeHandler(engine),
		},
	}
	return sv
}

func NewFakeEngine(accountKeyId string) *fake.Engine {
	return &fake.Engine{
		Clusters: []*objectStorageFakeApi.Cluster{
			{
				Id:              "isk01",
				ControlPanelUrl: "https://secure.sakura.ad.jp/objectstorage/",
				DisplayNameEnUs: "Ishikari Site #1",
				DisplayNameJa:   "石狩第1サイト",
				DisplayName:     "石狩第1サイト",
				DisplayOrder:    1,
				EndpointBase:    "/cloud/zone/is1a/api/objectstorage",
			},
		},
		Account: &objectStorageFakeApi.Account{
			Code:       "member@account@isk01",
			CreatedAt:  objectStorageFakeApi.CreatedAt(time.Now()),
			ResourceId: "100000000001",
		},
		AccountKeys: []*objectStorageFakeApi.AccountKey{{
			Id: objectStorageFakeApi.AccessKeyID(accountKeyId),
		}},
		Buckets: []*objectStorageFakeApi.Bucket{
			{
				ClusterId: "isk01",
				Name:      defaultBucketName,
			},
		},
		Permissions: []*objectStorageFakeApi.Permission{{
			Id:          objectStorageFakeApi.PermissionID(readonlyPermissionIdInt),
			DisplayName: objectStorageFakeApi.DisplayName(samplePermissionName),
		}},
		PermissionKeys: []*objectStorageFakeApi.PermissionKey{{
			Id: objectStorageFakeApi.AccessKeyID(readonlyPermissionKeyId),
		}},
	}
}
