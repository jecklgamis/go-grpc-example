default:
	cat ./Makefile

.PHONY: build
build: protos client server gateway

.PHONY: protos
protos:
	@protoc -I. \
	-I/usr/local/include \
	-I/usr/local/include/google/protobuf \
	-I$(GOPATH)/src \
	--go_out=plugins=grpc:. pkg/kvstore/kvstore.proto

	@protoc -I. \
	-I/usr/local/include \
	-I/usr/local/include/google/protobuf \
	-I$(GOPATH)/src \
	--grpc-gateway_out=logtostderr=true,grpc_api_configuration=pkg/kvstore/kvstore.yaml:. pkg/kvstore/kvstore.proto

get-protoc-gen-go:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: server
server:
	go build -o bin/server cmd/server/server.go
	@chmod +x bin/server

.PHONY: gateway
gateway:
	go build -o bin/gateway cmd/gateway/gateway.go
	@chmod +x bin/gateway

.PHONY: client
client:
	go build -o bin/client cmd/client/client.go
	@chmod +x bin/client

.PHONY: clean
clean:
	@rm -f bin/*
