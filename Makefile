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

DOCKERARGS := run -t --rm \
	-u $(shell id -u):$(shell id -g) \
	-v $(ROOTDIR):/usr/src \
	-e GOSUMDB=off \
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
generate: $(CACHEVOL) $(MODVOL) $(HOMEVOL)
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
docs: $(CACHEVOL) $(MODVOL) $(HOMEVOL)
	$(DOCKERENV) \
		protoc -I.:vendor:vendor/googleapis/:vendor/github.com/gogo/protobuf/protobuf/ \
			--doc_out=docs $(PROTOSOURCES) \
			--doc_opt=html,index-raw.html
	cat docs/header.txt docs/index-raw.html > docs/index.html
	rm docs/index-raw.html

# Generate API as typescript
.PHONY: ts
ts: $(CACHEVOL) $(MODVOL) $(HOMEVOL)
	@rm -Rf typescript
	@mkdir -p typescript
	$(DOCKERENV) \
		protoc -I.:vendor:vendor/googleapis/:vendor/github.com/gogo/protobuf/protobuf/ \
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

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: update-modules
update-modules:
	go get \
		github.com/golang/protobuf@v1.3.5

	go mod tidy
	go mod vendor

    # add .proto files manually
	cp -r vendor-proto/googleapis vendor/
	cp -r vendor-proto/github.com/gogo/protobuf/protobuf vendor/github.com/gogo/protobuf/
