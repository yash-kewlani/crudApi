package user

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yash-kewlani/crudApi/models"
	"github.com/yash-kewlani/crudApi/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var registerUserPayload models.RegisterUserRequest
	if err := utils.ParseJson(r, registerUserPayload); err != nil {
		log.Fatal("Error while parsing request body")
	}
}
