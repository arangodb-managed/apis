SHELL = bash
SCRIPTDIR := $(shell pwd)
ROOTDIR := $(shell cd $(SCRIPTDIR) && pwd)
BUILDIMAGE := arangodboasis/golang-ci:latest

PROTOSOURCES := $(shell find .  -name '*.proto' -not -path './vendor/*' -not -path './vendor-proto/*' | sort)

ifndef CIRCLECI
	GITHUB_TOKEN := $(shell cat $(HOME)/.arangodb/ms/github-readonly-code-acces.token)
else
	GITHUB_TOKEN :=
endif

GOMODCACHE := $(shell go env GOMODCACHE)
ifeq ($(GOMODCACHE),)
	GOMODCACHE := $(HOME)/go/pkg/mod
endif

DOCKERARGS := run -t --rm \
	-u $(shell id -u):$(shell id -g) \
	-v $(ROOTDIR):/usr/src \
	-v $(GOMODCACHE):/go/pkg/mod \
	-e GOMODCACHE=/go/pkg/mod \
	-e GOSUMDB=off \
	-e GOTOOLCHAIN=local \
	-e GOCACHE=/tmp/gocache \
	-e CGO_ENABLED=0 \
	-e GO111MODULE=on \
	-w /usr/src \
	$(BUILDIMAGE)

ifndef CIRCLECI
	DOCKERENV := docker $(DOCKERARGS)
else
	DOCKERENV :=
endif

.PHONY: all
all: generate build check ts docs

.PHONY: pull-build-image
pull-build-image: 
ifndef CIRCLECI
ifndef OFFLINE
	@docker pull $(BUILDIMAGE)
endif
endif

# Generate go code for proto files
.PHONY: generate
generate:
	$(DOCKERENV) \
		go generate ./...

# Build go code 
.PHONY: build
build: generate
	go build ./...

# Check go code 
.PHONY: check
check: 
	zutano go check ./...

# Generate API docs
.PHONY: docs
docs:
	$(DOCKERENV) \
		protoc -I.:vendor-proto/ \
			--doc_out=docs $(PROTOSOURCES) \
			--doc_opt=html,index-raw.html
	cat docs/header.txt docs/index-raw.html > docs/index.html
	rm docs/index-raw.html

# Generate API as typescript
.PHONY: ts
ts:
	@rm -Rf typescript
	@mkdir -p typescript
	$(DOCKERENV) \
		protoc -I.:vendor-proto \
			--ts_out=typescript $(PROTOSOURCES) \
			--ts_opt=.

.PHONY: test
test:
	mkdir -p bin/test
	go test -coverprofile=bin/test/coverage.out -v ./... | tee bin/test/test-output.txt ; exit "$${PIPESTATUS[0]}"
	cat bin/test/test-output.txt | go-junit-report > bin/test/unit-tests.xml
	go tool cover -html=bin/test/coverage.out -o bin/test/coverage.html

bootstrap:
	go get github.com/arangodb-managed/zutano
	go get github.com/jstemmer/go-junit-report
	go get github.com/stretchr/testify

check-version:
	zutano check api branch

.PHONY: update-modules
update-modules:
	go get \
		github.com/golang/protobuf@v1.5.4 \
		github.com/grpc-ecosystem/grpc-gateway@v1.16.0

	go mod tidy

.PHONY: update-vendor-proto
update-vendor-proto:
	@rm -rf vendor-proto
	@mkdir -p vendor-proto

	@# Copy google/api and google/rpc from grpc-gateway (which matches go.mod)
	@GRPC_GATEWAY_DIR=$$(go list -m -f '{{.Dir}}' github.com/grpc-ecosystem/grpc-gateway); \
	(cd $$GRPC_GATEWAY_DIR/third_party/googleapis && find google/api google/rpc -name "*.proto") | while read proto; do \
		mkdir -p vendor-proto/$$(dirname $$proto); \
		cp -f $$GRPC_GATEWAY_DIR/third_party/googleapis/$$proto vendor-proto/$$proto; \
	done

	@# Copy google/protobuf from protobuf repo (explicit version to avoid go.mod update)
	@go mod download github.com/protocolbuffers/protobuf@v3.20.3+incompatible
	@PROTOBUF_DIR=$$(go list -m -f '{{.Dir}}' github.com/protocolbuffers/protobuf@v3.20.3+incompatible); \
	(cd $$PROTOBUF_DIR/src && find google/protobuf -maxdepth 1 -name "*.proto" ! -name "*test*.proto") | while read proto; do \
		mkdir -p vendor-proto/$$(dirname $$proto); \
		cp -f $$PROTOBUF_DIR/src/$$proto vendor-proto/$$proto; \
		echo "Copied $$proto"; \
	done; \
	# Always include google/protobuf/compiler/plugin.proto \
	mkdir -p vendor-proto/google/protobuf/compiler; \
	cp -f $$PROTOBUF_DIR/src/google/protobuf/compiler/plugin.proto vendor-proto/google/protobuf/compiler/plugin.proto; \
	echo "Copied google/protobuf/compiler/plugin.proto";

	@chmod -R u+w vendor-proto
