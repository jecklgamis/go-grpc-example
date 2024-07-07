package main

import (
	"flag"
	"github.com/jecklgamis/go-grpc-example/pkg/gateway"
)

var (
	port           = flag.Int("port", 8080, "The gateway port")
	grpcServerAddr = flag.String("grpcServerAddr", "localhost:4000", "The server port")
)

func main() {
	flag.Parse()
	gateway.Start(*port, *grpcServerAddr)
}
