default:
	cat ./Makefile

.PHONY: build
build: protos client server

.PHONY: protos
protos:
	protoc --go_out=plugins=grpc:. pkg/kvstore/kvstore.proto

get-protoc-gen-go:
	go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: server
server:
	go build -o bin/server cmd/server/server.go
	chmod +x bin/server

.PHONY: client
client:
	go build -o bin/client cmd/client/client.go
	chmod +x bin/client

.PHONY: clean
clean:
	@rm -f bin/*
