# Copyright 2024 AboveCloud9.AI Products and Services Private Limited
# All rights reserved.
# This code may not be used, copied, modified, or distributed without explicit permission.

.PHONY: all
all: build
FORCE: ;

SHELL  := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= dev

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies build-coursed build-sfsyncd

build-coursed: 
	go build -tags $(LIBRARY_ENV) -o ./bin/coursed coursed/main.go

build-sfsyncd: 
	go build -tags $(LIBRARY_ENV) -o ./bin/sfsyncd sfsyncd/main.go

#build-cmd:
#	go build -tags $(LIBRARY_ENV) -o ./bin/search cmd/main.go

# Adding more flags for complete static linking
linux-binaries:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(LIBRARY_ENV) netgo" -installsuffix netgo -ldflags="-w -s" -o $(BIN_DIR)/coursed coursed/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(LIBRARY_ENV) netgo" -installsuffix netgo -ldflags="-w -s" -o $(BIN_DIR)/sfsyncd sfsyncd/main.go
#	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(LIBRARY_ENV) netgo" -installsuffix netgo -ldflags="-w -s" -o $(BIN_DIR)/search cmd/main.go

ci: dependencies test	

# builds one docker image
# TODO: extend for multiple services
docker:
	docker build -f Dockerfile.coursed -t coursed:$(LIBRARY_ENV) .
	docker build -f Dockerfile.sfsyncd -t sfsyncd:$(LIBRARY_ENV) .

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