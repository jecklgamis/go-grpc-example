## grpc-go-example

This is an example GRPC client, server, and gateway implementation. The server is a simple in-memory key-value store.

## Building

Install protocol buffers:

Mac OS X: 
```
brew install protobuf
```

Ensure version is 3.0.0 or above:
```
$ protoc --version
libprotoc 3.7.1
```

Install `protoc-gen-go` and related binaries
```
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```
or simple do `make get-protoc-gen-go`

Generate stub and build client, server, and gateway implementations:
```
make build
```

Example run:
```
$ make build
protoc --go_out=plugins=grpc:. pkg/kvstore/kvstore.proto
go build -o bin/client cmd/client/client.go
go build -o bin/server cmd/server/server.go
```

## Running 
Server:
```
$ bin/server -port 4000
```

Example output:
```
$ bin/server -port 4000
2019/08/13 07:30:24 Started server on port 4000
```

Gateway:
```
$ bin/gateway -port 8080 -grpcServerAddr localhost:4000
```
Example output:
```
2019/08/13 08:27:25 Using gRPC server on localhost:4000
2019/08/13 08:27:25 Started gateway on port 8080
```

Client:
```
$ bin/client -serverAddr localhost:4000
```

Example run:
```
bin/client -serverAddr localhost:4000
2019/08/12 07:50:19 Connecting to localhost:4000
2019/08/12 07:50:19 PUT ok some-key:f887bf33-bff5-4568-9d6d-9a7b329169da
2019/08/12 07:50:19 GET ok some-key:f887bf33-bff5-4568-9d6d-9a7b329169da
```

## Testing The REST EndPoints

```
curl -v -X PUT  http://localhost:8080/v1 -d@request_body.json
curl -v  http://localhost:8080/v1?key=some-key
```

request_body.json:
```
{
  "key": "some-key",
  "value": "some-value"
}
```


