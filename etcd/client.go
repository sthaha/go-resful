package etcd

import (
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Client struct {
	*clientv3.Client
}

func New() *Client {
	c, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		// Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"}
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		log.Fatalf("Unable to connect to etcd store: %v", err)
	}
	return &Client{Client: c}
}
