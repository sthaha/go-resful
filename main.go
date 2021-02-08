package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/sthaha/go-restful-example/userservice"
)

func main() {
	restful.Add(userservice.New())
	log.Fatal(http.ListenAndServe(":8080", nil))
}