package client

import (
	"context"
	"flag"
	pb "github.com/jecklgamis/grpc-go-example/pkg/kvstore"
	"google.golang.org/grpc"
	"log"
	"time"
)

type SomeClient struct {
	grpcClient pb.KeyValueStoreClient;
}

func New(serverAddr string) *SomeClient {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	grpcClient := pb.NewKeyValueStoreClient(conn)
	client := &SomeClient{grpcClient}
	return client

}

func (client *SomeClient) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := client.grpcClient.Get(ctx, &pb.Key{Key: key})
	if err != nil {
		log.Println("GET failed with", err)
		return "", err
	}
	log.Printf("GET ok %v:%v", key, response.Value)
	return response.Value, nil;
}

func (client *SomeClient) Put(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	kv := &pb.KeyValue{Key: &pb.Key{Key: key}, Value: &pb.Value{Value: value}};
	_, err := client.grpcClient.Put(ctx, kv)
	if err != nil {
		log.Println("PUT failed with", err)
		return err
	}
	log.Printf("PUT ok %v:%v", key, value)
	return nil;
}
