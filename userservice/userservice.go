package userservice

import (
	"encoding/json"
	"net/http"

	"context"
	"log"

	"github.com/emicklei/go-restful"

	client "github.com/sthaha/go-restful-example/etcd"
)

type User struct {
	ID        string `json:",omitempty"`
	FirstName string `json:",omitempty"`
	LastName  string `json:",omitempty"`
}

func GetUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	key := "/users/" + id

	res, err := client.Kv.Get(context.TODO(), key)
	if err != nil {
		log.Fatalf("%v", err)
	}

	usr := new(User)
	for _, ev := range res.Kvs {
		err := json.Unmarshal(ev.Value, &usr)
		if err != nil {
			response.WriteError(http.StatusInternalServerError, err)
		}
	}

	response.WriteEntity(usr)
}

func CreateUser(request *restful.Request, response *restful.Response) {
	usr := User{ID: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	usrData, err := json.Marshal(usr)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	_, err = client.Kv.Put(context.TODO(), "/users/"+usr.ID, string(usrData))
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	response.WriteEntity(usr)
}

func UpdateUser(request *restful.Request, response *restful.Response) {
	usr := User{ID: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	usrData, err := json.Marshal(usr)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	_, err = client.Kv.Put(context.TODO(), "/users/"+usr.ID, string(usrData))
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	response.WriteEntity(usr)
}

func DeleteUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	key := "/users/" + id

	res, err := client.Kv.Get(context.TODO(), key)
	if err != nil {
		log.Fatalf("%v", err)
	}

	usr := new(User)
	for _, ev := range res.Kvs {
		err := json.Unmarshal(ev.Value, &usr)
		if err != nil {
			response.WriteError(http.StatusInternalServerError, err)
		}
	}

	_, err = client.Kv.Delete(context.TODO(), "/users/"+usr.ID)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}
	response.WriteEntity(usr)
}

func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/users").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	service.Route(service.GET("/{user-id}").To(GetUser))
	service.Route(service.PUT("").To(UpdateUser))
	service.Route(service.POST("/{user-id}").To(CreateUser))
	service.Route(service.DELETE("/{user-id}").To(DeleteUser))
	return service
}
