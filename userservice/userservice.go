package userservice

import (
	"encoding/json"
	"net/http"
	"time"

	"context"
	"log"

	"github.com/emicklei/go-restful"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Client struct {
	*clientv3.Client
}

type User struct {
	ID        string `json:",omitempty"`
	FirstName string `json:",omitempty"`
	LastName  string `json:",omitempty"`
}

func (c *Client) GetUser(request *restful.Request, response *restful.Response) {
	kv := clientv3.NewKV(c.Client)

	id := request.PathParameter("user-id")
	key := "/users/" + id

	res, err := kv.Get(context.TODO(), key)
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

func (c *Client) CreateUser(request *restful.Request, response *restful.Response) {
	kv := clientv3.NewKV(c.Client)

	usr := User{ID: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	usrData, err := json.Marshal(usr)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	_, err = kv.Put(context.TODO(), "/users/"+usr.ID, string(usrData))
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	response.WriteEntity(usr)
}

func (c *Client) UpdateUser(request *restful.Request, response *restful.Response) {
	kv := clientv3.NewKV(c.Client)

	usr := User{}
	err := request.ReadEntity(&usr)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	usrData, err := json.Marshal(usr)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	_, err = kv.Put(context.TODO(), "/users/"+usr.ID, string(usrData))
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	response.WriteEntity(usr)
}

func (c *Client) DeleteUser(request *restful.Request, response *restful.Response) {
	kv := clientv3.NewKV(c.Client)

	id := request.PathParameter("user-id")
	key := "/users/" + id

	res, err := kv.Get(context.TODO(), key)
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

	_, err = kv.Delete(context.TODO(), "/users/"+usr.ID)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}
	response.WriteEntity(usr)
}

func New(c *Client) *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/users").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	service.Route(service.GET("/{user-id}").To(c.GetUser))
	service.Route(service.PUT("").To(c.UpdateUser))
	service.Route(service.POST("/{user-id}").To(c.CreateUser))
	service.Route(service.DELETE("/{user-id}").To(c.DeleteUser))
	return service
}

func EtcdClient() *Client {
	c, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		// Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"}
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		log.Fatalf("Unable to connect to etcd store: %v", err)
	}
	return &Client{Client: c}
}
