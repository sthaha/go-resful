package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/sthaha/go-restful-example/etcd"
	"github.com/sthaha/go-restful-example/services/user"
)

func main() {
	c, err := etcd.New()
	if err != nil {
		log.Fatalf("Unable to connect to etcd store: %v", err)
	}
	restful.Add(user.NewService(c))
	addr := ":8080"
	log.Print("Running at ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
