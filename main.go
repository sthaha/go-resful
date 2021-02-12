package main

import (

	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/sthaha/go-restful-example/etcd"
	"github.com/sthaha/go-restful-example/userservice"
)


func main() {
	c := etcd.New()
	restful.Add(userservice.New((*userservice.Client)(c)))
	addr := ":8080"
	log.Print("Running at ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

