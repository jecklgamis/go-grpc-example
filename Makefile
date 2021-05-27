default:
	cat ./Makefile

install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

.PHONY: build
build: protobufs client server  gateway-protobufs gateway

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
	@rm -f pkg/kvstore/*.go
