package models

type CustomerRegister struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}