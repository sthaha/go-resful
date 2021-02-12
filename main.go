package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/sthaha/go-restful-example/userservice"
)


func main() {
	c := userservice.EtcdClient()
	restful.Add(userservice.New(c))
	addr := ":8080"
	log.Print("Running at ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

