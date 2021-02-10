package main

import (
	"github.com/emicklei/go-restful"
	"github.com/sthaha/go-restful-example/userservice"
	"log"
	"net/http"
)

func main() {
	restful.Add(userservice.New())
	addr := ":8080"
	log.Print("Running at ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}