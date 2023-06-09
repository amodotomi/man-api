package user
// TODO | COMPLETELY CHANGE MODEL FOR WORKING WITH METHEOROLOGY DATA 
import (
	"fmt"
	"net/http"
	apperror "proj/internal/app-error"
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
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))
}


func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	// w.WriteHeader(200)
	// w.Write([]byte("get list of users"))

	// return nil
	return apperror.ErrNotFound
}           

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	// w.WriteHeader(201)
	// w.Write([]byte("create user"))

	// return nil
	return fmt.Errorf("API error")
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	// w.WriteHeader(200)
	// w.Write([]byte("get user by uuid"))

	// return nil
	return apperror.NewAppError(nil, "test", "test", "t13")
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("get list"))

	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("partially update user"))

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("delete user"))

	return nil
}
