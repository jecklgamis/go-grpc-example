package main

import (
	"flag"
	"github.com/google/uuid"
	"github.com/jecklgamis/grpc-go-example/pkg/client"
)

func String(str string) *string {
	return &str;
}

var (
	serverAddr = flag.String("serverAddr", "localhost:4000", "The server address")
)

func main() {
	client := client.New(*serverAddr)
	client.Put("some-key", uuid.New().String());
	client.Get("some-key")
}
