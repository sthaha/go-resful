package etcdclient

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func GetUserKV(id string) *clientv3.GetResponse {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	kv := clientv3.NewKV(cli)

	key := "/users/" + id

	result, err := kv.Get(context.TODO(), key)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return result
}

