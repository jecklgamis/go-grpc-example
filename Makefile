IMAGE_NAME:=go-grpc-example
IMAGE_TAG:=main

BUILD_BRANCH:=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_VERSION:=$(shell git rev-parse HEAD)
LD_FLAGS:="-X github.com/jecklgamis/go-grpc-example/pkg/version.BuildVersion=$(BUILD_VERSION) \
		  -X github.com/jecklgamis/go-grpc-example/pkg/version.BuildBranch=$(BUILD_BRANCH)"

default:
	@echo "Make targets:"
	@echo "make build - build binaries"
	@echo "make image - build Docker image"
	@echo "make run - run Docker image"

install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
update-deps:
	go get -u ./...

.PHONY: build
build: protobufs gateway-protobufs client server gateway

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

.PHONY: server
server: server-linux-amd64
	@echo "Building $@"
	@go build -o bin/$@ cmd/server/server.go
	@chmod +x bin/$@
server-linux-amd64:
	@echo "Building $@"
	@GOOS=linux GOARCH=amd64 go build -ldflags $(LD_FLAGS) -o bin/$@ cmd/server/server.go
	@chmod +x bin/$@
.PHONY: gateway
gateway: gateway-linux-amd64
	@echo "Building $@"
	@go build -o bin/$@ cmd/gateway/gateway.go
	@chmod +x bin/$@
gateway-linux-amd64:
	@echo "Building $@"
	@GOOS=linux GOARCH=amd64 go build -ldflags $(LD_FLAGS) -o bin/$@ cmd/gateway/gateway.go
	@chmod +x bin/$@

.PHONY: client
client: client-linux-amd64
	@echo "Building $@"
	go build -o bin/$@ cmd/client/client.go
	@chmod +x bin/$@
client-linux-amd64:
	@echo "Building $@"
	@GOOS=linux GOARCH=amd64 go build -ldflags $(LD_FLAGS) -o bin/$@ cmd/client/client.go
	@chmod +x bin/$@
.PHONY: clean
clean:
	@rm -f bin/*
	@rm -f pkg/kvstore/*.go
all: build image

image:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
run:
	docker run -p 4000:4000  -p 8080:8080 $(IMAGE_NAME):$(IMAGE_TAG)
run-bash:
	docker run -i -t $(IMAGE_NAME):$(IMAGE_TAG) /bin/bash
up:  all run