package models

type RegisterUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}
