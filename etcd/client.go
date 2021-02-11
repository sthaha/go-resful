package etcdclient

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func New() (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		// Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"}
		DialTimeout: 10 * time.Second,
	})
}
