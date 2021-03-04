package user

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sthaha/go-restful-example/app"

	"github.com/emicklei/go-restful"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type User struct {
	ID        string `json:",omitempty"`
	FirstName string `json:",omitempty"`
	LastName  string `json:",omitempty"`
}

type service struct {
	app app.App
}

func (s *service) Get(request *restful.Request, response *restful.Response) {
	kv := clientv3.NewKV(s.app.Etcd())

	id := request.PathParameter("user-id")
	key := "/users/" + id

	res, err := kv.Get(context.TODO(), key)
	if err != nil {
		log.Printf("err: %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	usr := new(User)
	for _, ev := range res.Kvs {
		err := json.Unmarshal(ev.Value, &usr)
		if err != nil {
			log.Printf("err: %v", err)
			response.WriteError(http.StatusInternalServerError, err)
			return
		}
	}

	response.WriteEntity(usr)
}

func (s *service) Create(request *restful.Request, response *restful.Response) {
	kv := clientv3.NewKV(s.app.Etcd())

	usr := User{ID: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	if err != nil {
		log.Printf("err: %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	usrData, err := json.Marshal(usr)
	if err != nil {
		log.Printf("err: %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	_, err = kv.Put(context.TODO(), "/users/"+usr.ID, string(usrData))
	if err != nil {
		log.Printf("err: %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	response.WriteEntity(usr)
}

func (s *service) Update(request *restful.Request, response *restful.Response) {
	kv := clientv3.NewKV(s.app.Etcd())

	usr := User{}
	err := request.ReadEntity(&usr)
	if err != nil {
		log.Printf("err: %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	usrData, err := json.Marshal(usr)
	if err != nil {
		log.Printf("err: %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	_, err = kv.Put(context.TODO(), "/users/"+usr.ID, string(usrData))
	if err != nil {
		log.Printf("err: %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	response.WriteEntity(usr)
}

func (s *service) Delete(request *restful.Request, response *restful.Response) {
	kv := clientv3.NewKV(s.app.Etcd())

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
			log.Printf("err: %v", err)
			response.WriteError(http.StatusInternalServerError, err)
			return
		}
	}

	_, err = kv.Delete(context.TODO(), "/users/"+usr.ID)
	if err != nil {
		log.Printf("err: %v", err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	response.WriteEntity(usr)
}

func NewService(a app.App) *restful.WebService {
	ws := &restful.WebService{}
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	user := &service{app: a}

	ws.Route(ws.GET("/{user-id}").To(user.Get))
	ws.Route(ws.PUT("").To(user.Update))
	ws.Route(ws.POST("/{user-id}").To(user.Create))
	ws.Route(ws.DELETE("/{user-id}").To(user.Delete))
	return ws
}
