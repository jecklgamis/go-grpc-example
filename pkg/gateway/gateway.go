package gateway

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jecklgamis/grpc-go-example/pkg/kvstore"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func Start(port int, grpcServerAddr string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := kvstore.RegisterKeyValueStoreHandlerFromEndpoint(ctx, mux, grpcServerAddr, opts)
	if err != nil {
		log.Fatal("Unable to register handler to multiplexer", err)
	}
	log.Println("Started gateway on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Fatal("Unable to start gateway", err)
	}
}
