package client

import (
	"context"
	"flag"
	pb "github.com/jecklgamis/grpc-go-example/pkg/kvstore"
	"google.golang.org/grpc"
	"log"
	"time"
)

type keyValueStoreClient struct {
	grpcClient pb.KeyValueStoreClient
}

func New(serverAddr string) *keyValueStoreClient {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	grpcClient := pb.NewKeyValueStoreClient(conn)
	client := &keyValueStoreClient{grpcClient}
	return client

}

func (client *keyValueStoreClient) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, _ := client.grpcClient.Get(ctx, &pb.Key{Key: key})
	if response != nil {
		return response.Value, nil
	}
	return "", nil
}

func (client *keyValueStoreClient) Put(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	kv := &pb.KeyValue{Key: key, Value: value}
	_, err := client.grpcClient.Put(ctx, kv)
	if err != nil {
		return err
	}
	return nil
}
