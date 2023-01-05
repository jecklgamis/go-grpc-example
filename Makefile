IMAGE_NAME:=jecklgamis/grpc-go-example
IMAGE_TAG:=$(shell git rev-parse HEAD)

BUILD_BRANCH:=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_VERSION:=$(shell git rev-parse HEAD)

default:
	cat ./Makefile

install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
update-modules:
	go get -u ./...
	go mod tidy

.PHONY: build
build: protobufs gateway-protobufs client server gateway

all: build

.PHONY: protobufs
protobufs:
	protoc -I . \
	--go_out=. --go_opt=paths=source_relative \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
           pkg/kvstore/kvstore.proto
gateway-protobufs:
	protoc -I. \
	-I$(GOPATH)/src \
	--grpc-gateway_out=logtostderr=true,grpc_api_configuration=pkg/kvstore/kvstore.yaml:. pkg/kvstore/kvstore.proto

LD_FLAGS:="-X github.com/jecklgamis/go-api-server-example/pkg/version.BuildVersion=$(BUILD_VERSION) \
		  -X github.com/jecklgamis/grpc-go-example/pkg/version.BuildBranch=$(BUILD_BRANCH)"

server: server-linux-amd64
	go build -o bin/server cmd/server/server.go
	@chmod +x bin/server

server-linux-amd64:
	@echo "Building $@"
	@GOOS=linux GOARCH=amd64 go build -ldflags $(LD_FLAGS) -o bin/server-linux-amd64 cmd/server/server.go
	@chmod +x bin/server-linux-amd64

gateway: gateway-linux-amd64
	go build -o bin/gateway cmd/gateway/gateway.go
	@chmod +x bin/gateway

gateway-linux-amd64:
	@echo "Building $@"
	@GOOS=linux GOARCH=amd64 go build -ldflags $(LD_FLAGS) -o bin/gateway-linux-amd64 cmd/gateway/gateway.go
	@chmod +x bin/gateway-linux-amd64

client: client-linux-amd64
	go build -o bin/client cmd/client/client.go
	@chmod +x bin/client

client-linux-amd64:
	@echo "Building $@"
	@GOOS=linux GOARCH=amd64 go build -ldflags $(LD_FLAGS) -o bin/client-linux-amd64 cmd/client/client.go
	@chmod +x bin/client-linux-amd64

clean:
	@rm -f bin/*
	@rm -f pkg/kvstore/*.go

up: build image run
image:
	@docker build -t $(IMAGE_NAME)/$(IMAGE_TAG) .
run:
	@docker run -p 4000:4000 -it $(IMAGE_NAME)/$(IMAGE_TAG)
run-shell:
	@docker run -i -t $(IMAGE_NAME)/$(IMAGE_TAG) /bin/bash
exec-shell:
	@docker exec -it `docker ps | grep $(IMAGE_NAME) | awk '{print $$1}'` /bin/bash
