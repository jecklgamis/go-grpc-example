## grpc-go-example

This is an example GRPC client and server. The server is a simple in-memory remote key-value store.

## Building
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

Example run:
```
$ bin/server -port 4000
2019/08/12 07:50:03 Started server on port 4000
2019/08/12 07:50:19 PUT key=some-key, value=f887bf33-bff5-4568-9d6d-9a7b329169da
2019/08/12 07:50:19 GET key=some-key, value=f887bf33-bff5-4568-9d6d-9a7b329169da
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

