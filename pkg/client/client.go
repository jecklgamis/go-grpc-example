package client

import (
	"context"
	"crypto/tls"
	"flag"
	pb "github.com/jecklgamis/go-grpc-example/pkg/kvstore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type keyValueStoreClient struct {
	grpcClient pb.KeyValueStoreClient
}

func New(serverAddr string, ssl bool) *keyValueStoreClient {
	flag.Parse()
	var opts []grpc.DialOption
	if ssl {
		config := &tls.Config{
			InsecureSkipVerify: true,
		}
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
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
