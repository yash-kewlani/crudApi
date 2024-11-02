package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/yash-kewlani/crudApi/models"
	"github.com/yash-kewlani/crudApi/service/auth"
	"github.com/yash-kewlani/crudApi/utils"
)

type Handler struct {
	store models.UserStore
}

func NewHandler(store models.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var registerUserPayload models.RegisterUserRequest
	if err := utils.ParseJson(r, &registerUserPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(registerUserPayload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request body %v", errors))
		return
	}

	_, err := h.store.GetUserByEmail(registerUserPayload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", registerUserPayload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(registerUserPayload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(models.User{
		FirstName: registerUserPayload.FirstName,
		LastName:  registerUserPayload.LastName,
		Email:     registerUserPayload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusCreated, nil)
}
