package server

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net"
)
import pb "github.com/jecklgamis/grpc-go-example/pkg/kvstore"
import "google.golang.org/grpc/status"

type keyValueStoreServer struct {
	store map[string]string;
}

var server = &keyValueStoreServer{store: make(map[string]string)}

func (server *keyValueStoreServer) Get(ctx context.Context, key *pb.Key) (*pb.Value, error) {
	if key == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid key")
	}
	value := server.store[key.Key]
	log.Printf("GET key=%v, value=%v\n", key.Key, value)
	return &pb.Value{Value: value}, nil
}

func (server *keyValueStoreServer) Put(ctx context.Context, kv *pb.KeyValue) (*pb.Empty, error) {
	log.Printf("PUT key=%v, value=%v\n", kv.Key.Key, kv.Value.Value)
	server.store[kv.Key.Key] = kv.Value.Value
	return &pb.Empty{}, nil
}

func Start(port int) {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterKeyValueStoreServer(grpcServer, server)
	go func() {
		log.Println("Started server on port", port)
	}();
	grpcServer.Serve(listener)
}
