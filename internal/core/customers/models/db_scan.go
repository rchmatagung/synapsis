package models

type Customer struct {
	CustomerId int64  `json:"customer_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	UpdatedAt  string `json:"updated_at"`
	UpdatedBy  string `json:"updated_by"`
	DeletedAt  string `json:"deleted_at"`
	DeletedBy  string `json:"deleted_by"`
}