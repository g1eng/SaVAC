NAME = savac
VARIANTS = o12 usaco sadns
MODEL_PATH = pkg/models
DOCKER := docker
GOLANG_CI_LINT_VERSION := v2.1.6

.PHONY: all build install install_all download_spec release mock_server test coverage clean

all: build

build:
	cd cmd && go build -o ../$(NAME)
	strip $(NAME)
install: release
	[ -d $(PREFIX)/bin ] || mkdir -vp $(PREFIX)/bin 
	install -m 755 $(NAME) $(PREFIX)/bin/$(NAME)
install_all: install
	@$(foreach v,$(VARIANTS), \
		install -m 755 $(NAME) $(PREFIX)/bin/$(v);\
	)

download_spec:
	[ -d spec ] || mkdir spec \
	&& curl -fSL -o spec/openapi.json https://manual.sakura.ad.jp/vps/api/api-doc/api-json.json

test: prepare_minio
	go test -test.v -p 8 ./... && make teardown_minio || make teardown_minio && exit 1

acc_test:
	TESTACC=1 go test -test.v -p 2 ./...

tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin $(GOLANG_CI_LINT_VERSION)

prepare_minio:
	$(DOCKER) run  -d --rm --name min1 -p 19002:19002 -p 19003:19003  quay.io/minio/minio server /data --address ":19002" --console-address ":19003" \
	&& $(DOCKER) run  -d --rm --name min2 -p 29002:29002 -p 29003:29003  quay.io/minio/minio server /data --address ":29002" --console-address ":29003"

teardown_minio:
	$(DOCKER) stop min1 \
	&& $(DOCKER) stop min2

# E2E test is out of scope for CI/CD at now
coverage:
	go test \
		-v \
		-coverprofile=coverage.txt \
		-timeout 0 \
		-covermode=atomic \
		-coverpkg  ./cmd/...,./pkg/vps/...,./pkg/cloud/sacloud/... \
		-p 8 \
		./cmd/... \
		./pkg/vps/... \
		./pkg/cloud/sacloud/...

lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run --fix ./...

e2e_test:
	TESTACC=1 go test -v -timeout 0 -p 1 ./cmd/...

mock_server:
	cd test_utils/mock_server \
	&& go build -o mock_server

release: build
	strip $(NAME)

clean:
	rm -v $(NAME)
