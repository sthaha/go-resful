package client

import (
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var EtcdClient *clientv3.Client

func init() {
	var err error
	// create connection to etcd store
	EtcdClient, err = New()
	if err != nil {
		log.Fatalf("Unable to connect to etcd store: %v", err)
	}
}

func New() (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		// Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"}
		DialTimeout: 10 * time.Second,
	})
}
