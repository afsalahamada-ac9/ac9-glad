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

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies build-coursed build-sfsyncd build-mediad

build-coursed: 
	go build -tags $(TAG) -o ./bin/coursed $(SERVICES_DIR)/coursed/main.go

build-sfsyncd: 
	go build -tags $(TAG) -o ./bin/sfsyncd $(SERVICES_DIR)/sfsyncd/main.go

build-mediad:
	go build -tags $(TAG) -o ./bin/mediad $(SERVICES_DIR)/mediad/main.go

#build-cmd:
#	go build -tags $(TAG) -o ./bin/search cmd/main.go

# Adding more flags for complete static linking
linux-binaries:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(TAG) netgo" -installsuffix netgo -ldflags="-w -s" -o $(BIN_DIR)/coursed $(SERVICES_DIR)/coursed/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(TAG) netgo" -installsuffix netgo -ldflags="-w -s" -o $(BIN_DIR)/sfsyncd $(SERVICES_DIR)/sfsyncd/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(TAG) netgo" -installsuffix netgo -ldflags="-w -s" -o $(BIN_DIR)/mediad $(SERVICES_DIR)/mediad/main.go
#	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(TAG) netgo" -installsuffix netgo -ldflags="-w -s" -o $(BIN_DIR)/search cmd/main.go

ci: dependencies test	

# builds one docker image
# TODO: support easy extension for multiple services
docker:
	docker build -f $(DOCKERFILE_DIR)/Dockerfile.coursed -t coursed:$(TAG) .
	docker build -f $(DOCKERFILE_DIR)/Dockerfile.sfsyncd -t sfsyncd:$(TAG) .
	docker build -f $(DOCKERFILE_DIR)/Dockerfile.mediad -t mediad:$(TAG) .

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