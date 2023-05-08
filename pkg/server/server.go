package server

import (
	"context"
	"fmt"
	pb "github.com/jecklgamis/grpc-go-example/pkg/kvstore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type keyValueStoreServer struct {
	pb.UnimplementedKeyValueStoreServer
	store map[string]string
}

var server = &keyValueStoreServer{store: make(map[string]string)}

func (server *keyValueStoreServer) Get(ctx context.Context, key *pb.Key) (*pb.Value, error) {
	if key == nil {
		log.Println("GET: Invalid key")
		return nil, status.Errorf(codes.InvalidArgument, "Invalid key")
	}
	v, ok := server.store[key.Key]
	if !ok {
		log.Printf("GET : Key not found : %v\n", key.Key)
		return nil, status.Errorf(codes.NotFound, "Key not found")
	}
	log.Printf("GET : key=%v, value=%v\n", key.Key, v)
	return &pb.Value{Value: v}, nil
}

func (server *keyValueStoreServer) Put(ctx context.Context, kv *pb.KeyValue) (*pb.Empty, error) {
	if kv == nil {
		return &pb.Empty{}, status.Errorf(codes.InvalidArgument, "Invalid key value")
	}
	if kv.Key == "" {
		return &pb.Empty{}, status.Errorf(codes.InvalidArgument, "Empty key not allowed")
	}
	if kv.Value == "" {
		return &pb.Empty{}, status.Errorf(codes.InvalidArgument, "Empty value not allowed")
	}
	log.Printf("PUT : key=%v, value=%v\n", kv.Key, kv.Value)
	server.store[kv.Key] = kv.Value
	return &pb.Empty{}, nil
}

func Start(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterKeyValueStoreServer(grpcServer, server)
	log.Println("Started server on port", port)
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Unable to start server: %v", err)
	}
}
