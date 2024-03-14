package models

type ShoppingCart struct {
	ShoppingCartId       int64                 `json:"shopping_cart_id"`
	CustomerId           int64                 `json:"customer_id"`
	StatusCode           string                `json:"status_code"`
	PaymentTotal         int64                 `json:"payment_total"`
	ShoppingCardProducts []ShoppingCartProduct `json:"shopping_cart_products" gorm:"foreignKey:ShoppingCartId;references:ShoppingCartId"`
	CreatedAt            string                `json:"created_at"`
	CreatedBy            string                `json:"created_by"`
	UpdatedAt            string                `json:"updated_at"`
	UpdatedBy            string                `json:"updated_by"`
}

type ShoppingCartProduct struct {
	ShoppingCartProductId int64  `json:"shopping_cart_product_id"`
	ShoppingCartId        int64  `json:"shopping_cart_id"`
	ProductId             int64  `json:"product_id"`
	ProductName           string `json:"product_name"`
	Price                 int64  `json:"price"`
	Quantity              int64  `json:"quantity"`
	CategoryName          string `json:"category_name"`
	SubTotal              int64  `json:"sub_total"`
	CreatedAt             string `json:"created_at"`
	CreatedBy             string `json:"created_by"`
	UpdatedAt             string `json:"updated_at"`
	UpdatedBy             string `json:"updated_by"`
}
