package models

type Product struct {
	ProductId    int64  `json:"product_id"`
	ProductCode  string `json:"product_code"`
	ProductName  string `json:"product_name"`
	IsActive     bool   `json:"is_active"`
	Price        int64  `json:"price"`
	CategoryId   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
	Stock        int64  `json:"stock"`
	CreatedAt    string `json:"created_at"`
	CreatedBy    string `json:"created_by"`
	UpdatedAt    string `json:"updated_at"`
	UpdatedBy    string `json:"updated_by"`
}