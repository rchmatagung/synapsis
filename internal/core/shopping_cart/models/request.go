package models

type ShoppingCartRequest struct {
	StatusCode           string                       `json:"status_code"`
	PaymentTotal         int64                        `json:"payment_total"`
	ShoppingCardProducts []ShoppingCartProductRequest `json:"shopping_cart_products"`
}

type ShoppingCartProductRequest struct {
	ProductId int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
	SubTotal  int64 `json:"sub_total"`
	Price     int64 `json:"price"`
}