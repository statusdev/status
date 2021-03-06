# go settings
GOFLAGS := -mod=vendor
GO := GOFLAGS=$(GOFLAGS) GO111MODULE=on CGO_ENABLED=0 go
GOTEST := GOFLAGS=$(GOFLAGS) GO111MODULE=on CGO_ENABLED=1 go # -race needs cgo

ifndef DATE
	DATE := $(shell date -u '+%Y%m%d')
endif

ifndef SHA
	SHA := $(shell git rev-parse --short HEAD)
endif

.PHONY: apiv1
apiv1: api/v1/models api/v1/restapi

GOSWAGGER ?= docker run --rm \
	--user=$(shell id -u $(USER)):$(shell id -g $(USER)) \
	-v $(shell pwd):/go/src/github.com/statusdev/status \
	-w /go/src/github.com/statusdev/status quay.io/goswagger/swagger:v0.19.0

api/v1/models api/v1/restapi: swagger.yaml
	-rm -r api/v1/{models,restapi}
	$(GOSWAGGER) generate server -f swagger.yaml --exclude-main -A statusapi --target api/v1

SWAGGER ?= docker run --rm \
		--user=$(shell id -u $(USER)):$(shell id -g $(USER)) \
		-v $(shell pwd):/local \
		swaggerapi/swagger-codegen-cli:2.4.0

.PHONY: lint
lint:
	golint $(shell $(GO) list ./...)

.PHONY: check-vendor
check-vendor:
	$(GO) mod tidy
	$(GO) mod vendor
	git update-index --refresh
	git diff-index --quiet HEAD

.PHONY: test
test:
	$(GOTEST) test -coverprofile coverage.out -race -v ./...

.PHONY: build
build: cmd/api/api

.PHONY: cmd/api/api
cmd/api/api:
	$(GO) build -v -o ./cmd/api/api ./cmd/api

.PHONY: container-api
container-api: cmd/api/api
	docker build -t status ./cmd/api