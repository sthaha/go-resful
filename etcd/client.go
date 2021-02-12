package client

import (
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var Kv clientv3.KV

func init() {
	// create connection to etcd store
	etcdClient, err := New()
	if err != nil {
		log.Fatalf("Unable to connect to etcd store: %v", err)
	}
	clientv3.NewKV(etcdClient)
	Kv = clientv3.NewKV(etcdClient)
}

func New() (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		// Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"}
		DialTimeout: 10 * time.Second,
	})
}
