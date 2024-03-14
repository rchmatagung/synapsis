package models

type LoginResponse struct {
	CustomerId int64  `json:"customer_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}
