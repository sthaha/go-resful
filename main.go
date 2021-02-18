package main

import (
	"github.com/sthaha/go-restful-example/app"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/sthaha/go-restful-example/services/user"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatalf("Unable to initialize the app: %v", err)
	}
	restful.Add(user.NewService(a))
	addr := ":8080"
	log.Print("Running at ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
