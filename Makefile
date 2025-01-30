FIPS_ENABLED?=true

ifeq ($(GO),)
GO:=$(shell command -v go)
endif

GOHOSTOS:=$(shell $(GO) env GOHOSTOS)
GOPATH:=$(shell $(GO) env GOPATH)
GOOS?=$(shell $(GO) env GOOS)
GOARCH?=$(shell $(GO) env GOARCH)
GOBIN?=$(shell $(GO) env GOBIN)
GOFLAGS_MOD ?=

GOENV=GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=1 GOFLAGS="${GOFLAGS_MOD}"
GOBUILDFLAGS=-gcflags="all=-trimpath=${GOPATH}" -asmflags="all=-trimpath=${GOPATH}"

ifeq (${FIPS_ENABLED}, true)
GOFLAGS_MOD+=-tags=fips_enabled
GOFLAGS_MOD:=$(strip ${GOFLAGS_MOD})
GOENV+=GOEXPERIMENT=strictfipsruntime,boringcrypto
GOENV:=$(strip ${GOENV})
endif

IMAGE ?="quay.io/cloudservices/kessel-inventory"
IMAGE_TAG=$(git rev-parse --short=7 HEAD)
GIT_COMMIT=$(git rev-parse --short HEAD)

ifeq ($(DOCKER),)
DOCKER:=$(shell command -v podman || command -v docker)
endif

API_PROTO_FILES:=$(shell find api -name *.proto)

TITLE:="Kessel Asset Inventory API"
ifeq ($(VERSION),)
VERSION:=$(shell git describe --tags --always)
endif
INVENTORY_SCHEMA_VERSION=0.11.0

.PHONY: init
# init env
init:
	$(GO) install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	$(GO) install github.com/google/wire/cmd/wire@latest
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	$(GO) get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	$(GO) install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	$(GO) install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	$(GO) get github.com/envoyproxy/protoc-gen-validate
	$(GO) install github.com/envoyproxy/protoc-gen-validate

.PHONY: api
# generate api proto
api:
	@echo "Generating api protos"
	@$(DOCKER) build -t custom-protoc ./api
	@$(DOCKER) run -t --rm -v $(PWD)/api:/api:rw,z -v $(PWD)/openapi.yaml:/openapi.yaml:rw,z \
	-w=/api/ custom-protoc sh -c "buf generate && \
		buf lint && \
		buf breaking --against 'buf.build/project-kessel/inventory-api' "

.PHONY: api_breaking
# generate api proto
api_breaking:
	@echo "Generating api protos, allowing breaking changes"
	@$(DOCKER) build -t custom-protoc ./api
	@$(DOCKER) run -t --rm -v $(PWD)/api:/api:rw,z -v $(PWD)/openapi.yaml:/openapi.yaml:rw,z \
	-w=/api/ custom-protoc sh -c "buf generate && \
		buf lint"

# .PHONY: api
# # generate api proto
# api:
# 	@echo "Generating api protos"
# 	@$(DOCKER) build -t custom-protoc ./api
# 	@$(DOCKER) run -t --rm -v $(PWD)/api:/api:rw -v $(PWD)/openapi.yaml:/openapi.yaml:rw \
# 	-w=/api/ custom-protoc sh -c "buf generate && buf lint"

.PHONY: build
# build
build:
	$(warning Setting GOEXPERIMENT=strictfipsruntime,boringcrypto - this generally causes builds to fail unless building inside the provided Dockerfile. If building locally, run `make local-build`)
	mkdir -p bin/ && ${GOENV} GOOS=${GOOS} ${GO} build ${GOBUILDFLAGS} -ldflags "-X cmd.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: local-build
# local-build to ensure FIPS is not enabled which would likely result in a failed build locally
local-build:
	mkdir -p bin/ && $(GO) build -ldflags "-X cmd.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: docker-build-push
docker-build-push:
	./build_deploy.sh

.PHONY: build-push-minimal
build-push-minimal:
	./build_push_minimal.sh

.PHONY: clean
# removes all binaries
clean:
	rm -rf bin/

.PHONY: test
# run all tests
test:
	@echo ""
	@echo "Running tests."
	# TODO: e2e tests are taking too long to be enabled by default. They need to be sped up.
	@$(GO) test ./... -count=1 -coverprofile=coverage.out -skip 'TestInventoryAPIGRPC_*|TestInventoryAPIHTTP_*|Test_ACMKafkaConsumer'
	@echo "Overall test coverage:"
	@$(GO) tool cover -func=coverage.out | grep total: | awk '{print $$3}'

test-coverage: test
	@$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "coverage report written to coverage.html"


.PHONY: generate
# generate
generate:
	$(GO) mod tidy
	$(GO) get github.com/google/wire/cmd/wire@latest
	$(GO) generate ./...

.PHONY: all
# generate all
all:
	make api;
	# make config;
	make generate;

.PHONY: lint
# run go linter with the repositories lint config
lint:
	@echo "Linting code."
	@$(DOCKER) run -t --rm -v $(PWD):/app -w /app golangci/golangci-lint golangci-lint run -v

.PHONY: pr-check
# generate pr-check
pr-check:
	make generate;
	make test;
	make lint;
	make local-build;
	#

.PHONY: inventory-up
inventory-up:
	./scripts/start-inventory.sh

.PHONY: inventory-up-sso
inventory-up-sso:
	./scripts/start-inventory-kc.sh

.PHONY: inventory-up-kafka
inventory-up-kafka:
	./scripts/start-inventory-kafka.sh

.PHONY: inventory-up-kind
inventory-up-kind:
	./scripts/start-inventory-kind.sh

.PHONY: get-token
get-token:
	./scripts/get-token.sh

.PHONY: inventory-down
inventory-down:
	./scripts/stop-inventory.sh

.PHONY: inventory-down-sso
inventory-down-sso:
	./scripts/stop-inventory-kc.sh

.PHONY: inventory-down-kafka
inventory-down-kafka:
	./scripts/stop-inventory-kafka.sh

.PHONY: inventory-down-kind
inventory-down-kind:
	./scripts/stop-inventory-kind.sh

.PHONY: run
# run api locally
run: local-build
	$(GO) run main.go serve

run-help: local-build
	$(GO) run main.go serve --help

.PHONY: migrate
# run database migrations
migrate: local-build
	./bin/inventory-api migrate --config .inventory-api.yaml

.PHONY: update-schema
# fetch the latest schema from github.com/RedHatInsights/kessel-config
update-schema:
	./scripts/get-schema-yaml.sh > ./deploy/schema.yaml

help:
# show help
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
