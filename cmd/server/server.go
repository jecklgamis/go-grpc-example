package main

import (
	"flag"
	"fmt"
	"github.com/jecklgamis/grpc-go-example/pkg/gateway"
	"github.com/jecklgamis/grpc-go-example/pkg/server"
	"time"
)

var (
	port        = flag.Int("port", 4000, "The server port")
	gatewayPort = flag.Int("gatewayPort", 8080, "The gateway port")
)

func main() {
	flag.Parse()
	go server.Start(*port)
	time.Sleep(3 * time.Second)
	go gateway.Start(*gatewayPort, fmt.Sprintf(":%d", *port))
	for {
		time.Sleep(1 * time.Second)
	}
}
