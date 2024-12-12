# Copyright 2024 AboveCloud9.AI Products and Services Private Limited
# All rights reserved.
# This code may not be used, copied, modified, or distributed without explicit permission.

.PHONY: all
all: build
FORCE: ;

SHELL  := env TAG=$(TAG) $(SHELL)
TAG ?= dev
DOCKERFILE_DIR = ops/docker
SERVICES_DIR = services

GIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date +%Y-%m-%d_%H:%M:%S)
VERSION := $(shell cat .version)

BIN_DIR = $(PWD)/bin
LD_FLAGS = -ldflags="-w -s -X 'ac9/glad/pkg/util.gitHash=$(GIT_HASH)' \
	-X 'ac9/glad/pkg/util.buildTime=$(BUILD_TIME)' \
	-X 'ac9/glad/pkg/util.version=$(VERSION)'"
GO_BUILD_LINUX = CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(TAG) netgo" -installsuffix netgo 
GO_BUILD = go build -tags $(TAG)

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies build-coursed build-sfsyncd build-mediad build-ldsd build-gcd build-pushd

build-coursed: 
	$(GO_BUILD) $(LD_FLAGS) -o ./bin/coursed $(SERVICES_DIR)/coursed/main.go

build-sfsyncd: 
	$(GO_BUILD) $(LD_FLAGS) -o ./bin/sfsyncd $(SERVICES_DIR)/sfsyncd/main.go

build-mediad:
	$(GO_BUILD) $(LD_FLAGS) -o ./bin/mediad $(SERVICES_DIR)/mediad/main.go

build-ldsd:
	$(GO_BUILD) $(LD_FLAGS) -o ./bin/ldsd $(SERVICES_DIR)/ldsd/main.go

build-gcd:
	$(GO_BUILD) $(LD_FLAGS) -o ./bin/gcd $(SERVICES_DIR)/gcd/main.go

build-pushd:
	$(GO_BUILD) $(LD_FLAGS) -o ./bin/pushd $(SERVICES_DIR)/pushd/main.go

#build-cmd:
#	$(GO_BUILD) $(LD_FLAGS) -o ./bin/search cmd/main.go

# Adding more flags for complete static linking
linux-binaries:
	$(GO_BUILD_LINUX) $(LD_FLAGS) -o $(BIN_DIR)/coursed $(SERVICES_DIR)/coursed/main.go
	$(GO_BUILD_LINUX) $(LD_FLAGS) -o $(BIN_DIR)/sfsyncd $(SERVICES_DIR)/sfsyncd/main.go
	$(GO_BUILD_LINUX) $(LD_FLAGS) -o $(BIN_DIR)/mediad $(SERVICES_DIR)/mediad/main.go
	$(GO_BUILD_LINUX) $(LD_FLAGS) -o $(BIN_DIR)/ldsd $(SERVICES_DIR)/ldsd/main.go
	$(GO_BUILD_LINUX) $(LD_FLAGS) -o $(BIN_DIR)/gcd $(SERVICES_DIR)/gcd/main.go
	$(GO_BUILD_LINUX) $(LD_FLAGS) -o $(BIN_DIR)/pushd $(SERVICES_DIR)/pushd/main.go
#	$(GO_BUILD_LINUX) $(LD_FLAGS) -o $(BIN_DIR)/search cmd/main.go

ci: dependencies test	

# builds one docker image
# TODO: support easy extension for multiple services
docker:
	docker build -f $(DOCKERFILE_DIR)/Dockerfile.coursed -t coursed:$(TAG) .
	docker build -f $(DOCKERFILE_DIR)/Dockerfile.coursed -t coursed:$(TAG) .
	docker build -f $(DOCKERFILE_DIR)/Dockerfile.sfsyncd -t sfsyncd:$(TAG) .
	docker build -f $(DOCKERFILE_DIR)/Dockerfile.mediad -t mediad:$(TAG) .
	docker build -f $(DOCKERFILE_DIR)/Dockerfile.ldsd -t ldsd:$(TAG) .
	docker build -f $(DOCKERFILE_DIR)/Dockerfile.gcd -t gcd:$(TAG) .
	docker build -f $(DOCKERFILE_DIR)/Dockerfile.pushd -t pushd:$(TAG) .

build-mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=usecase/account/interface.go -destination=usecase/account/mock/account.go
	@~/go/bin/mockgen -source=usecase/center/interface.go -destination=usecase/center/mock/center.go
	@~/go/bin/mockgen -source=usecase/course/interface.go -destination=usecase/course/mock/course.go
	@~/go/bin/mockgen -source=usecase/product/interface.go -destination=usecase/product/mock/product.go
	@~/go/bin/mockgen -source=usecase/tenant/interface.go -destination=usecase/tenant/mock/tenant.go

test:
	go test -tags testing ./...

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done