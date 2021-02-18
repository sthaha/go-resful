package app

import (
	"github.com/sthaha/go-restful-example/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type App struct {
	Etcd *clientv3.Client
}

func New() (*App, error) {
	a := App{}

	// Initialize Etcd Client
	client, err := etcd.New()
	if err != nil {
		return nil, err
	}
	a.Etcd = client

	return &a, nil
}