package userservice

import (
	"github.com/emicklei/go-restful"
	"net/http"
)

type User struct {
	Name string
	ID string
}

func GetUser(request *restful.Request, response *restful.Response) {
	// some user := fetch by userid
	id := request.PathParameter("user-id")
	usr := &User{ID: id, Name: "John Doe"}
	response.WriteEntity(usr)
}

func UpdateUser(request *restful.Request, response *restful.Response) {
	// update user where user = userid
	usr := new(User)
	err := request.ReadEntity(&usr)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}
	response.WriteEntity(usr)
}

func CreateUser(request *restful.Request, response *restful.Response) {
	// new user id = userid
	usr := User{ID: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}
	response.WriteEntity(usr)
}

func DeleteUser(request *restful.Request, response *restful.Response) {
	// delete user where userid = userid
}

func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	service.Route(service.GET("/{user-id}").To(GetUser))
	service.Route(service.POST("").To(UpdateUser))
	service.Route(service.PUT("/{user-id}").To(CreateUser))
	service.Route(service.DELETE("/{user-id}").To(DeleteUser))

	return service
}
