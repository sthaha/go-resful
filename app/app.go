package app

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type App interface {
	Etcd() *clientv3.Client
}

type app struct {
	etcd *clientv3.Client
}

// ensure app is conforms to App interface
var _ App = (*app)(nil)

func New() (App, error) {
	// Initialize Etcd Client
	etcd, err := newEtcd()
	if err != nil {
		return nil, err
	}

	return &app{etcd: etcd}, nil
}

func (a *app) Etcd() *clientv3.Client {
	return a.etcd
}

func newEtcd() (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		// Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"}
		DialTimeout: 10 * time.Second,
	})
}
