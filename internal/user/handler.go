package user

import (
	"net/http"
	"proj/internal/handlers"
	"proj/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users" 
	userURL = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserByUUID)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}


func (h *handler) GetList( 
		w http.ResponseWriter,
		r *http.Request,
		params httprouter.Params,
	) {
	w.Write([]byte("get list of users"))
}           


func (h *handler) CreateUser(
		w http.ResponseWriter,
		r *http.Request,
		params httprouter.Params) {
	w.Write([]byte("create user"))
}

func (h *handler) GetUserByUUID(
		w http.ResponseWriter, 
		r *http.Request, 
		params httprouter.Params) {
	w.Write([]byte("get user by uuid")) 
}

func (h *handler) UpdateUser(
		w http.ResponseWriter, 
		r *http.Request, 
		params httprouter.Params) {
	w.Write([]byte("get list"))
}

func (h *handler) PartiallyUpdateUser(
		w http.ResponseWriter, 
		r *http.Request, 
		params httprouter.Params) {
	w.Write([]byte("partially update user"))
}

func (h *handler) DeleteUser(
		w http.ResponseWriter, 
		r *http.Request, 
		params httprouter.Params) {
	w.Write([]byte("delete user"))
}
