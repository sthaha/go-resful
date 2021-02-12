package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/sthaha/go-restful-example/userservice"
)


func main() {
	restful.Add(userservice.New())
	addr := ":8080"
	log.Print("Running at ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

