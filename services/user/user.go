package user

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"

	etcd "go.etcd.io/etcd/client/v3"
)

type User struct {
	ID        string `json:",omitempty"`
	FirstName string `json:",omitempty"`
	LastName  string `json:",omitempty"`
}

type service struct {
	etcd *etcd.Client
}

func (s *service) Get(request *restful.Request, response *restful.Response) {
	kv := etcd.NewKV(s.etcd)

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

func (s *service) Create(request *restful.Request, response *restful.Response) {
	kv := etcd.NewKV(s.etcd)

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

func (s *service) Update(request *restful.Request, response *restful.Response) {
	kv := etcd.NewKV(s.etcd)

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

func (s *service) Delete(request *restful.Request, response *restful.Response) {
	kv := etcd.NewKV(s.etcd)

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

func NewService(c *etcd.Client) *restful.WebService {
	ws := &restful.WebService{}
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	user := &service{c}

	ws.Route(ws.GET("/{user-id}").To(user.Get))
	ws.Route(ws.PUT("").To(user.Update))
	ws.Route(ws.POST("/{user-id}").To(user.Create))
	ws.Route(ws.DELETE("/{user-id}").To(user.Delete))
	return ws
}
