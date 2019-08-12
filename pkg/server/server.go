package server

import (
	"context"
	"fmt"
	pb "github.com/jecklgamis/grpc-go-example/pkg/kvstore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type keyValueStoreServer struct {
	store map[string]string
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
	if kv == nil {
		log.Printf("kv is nil")
		return &pb.Empty{}, nil
	}
	log.Printf("PUT key=%v, value=%v\n", kv.Key, kv.Value)
	server.store[kv.Key] = kv.Value
	return &pb.Empty{}, nil
}

func Start(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterKeyValueStoreServer(grpcServer, server)
	log.Println("Started server on port", port)
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Unable to start server: %v", err)
	}
}
