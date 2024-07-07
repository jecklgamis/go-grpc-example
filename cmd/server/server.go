package main

import (
	"flag"
	"github.com/jecklgamis/go-grpc-example/pkg/server"
)

var (
	port = flag.Int("port", 4000, "The server port")
)

func main() {
	flag.Parse()
	server.Start(*port)
}
