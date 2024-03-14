package shopping_cart

import (
	"synapsis/internal/middleware"
	"synapsis/internal/wrapper/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Post("/shopping_cart/addproduct", middleware.EmployeeAuthMiddleware(), handler.Core.Shopping_cart.AddProductShoppingCart)
	api.Get("/shopping_cart/get", middleware.EmployeeAuthMiddleware(), handler.Core.Shopping_cart.GetShoppingCart)
	api.Delete("/shopping_cart/deleteproduct/:shopping_cart_product_id", middleware.EmployeeAuthMiddleware(), handler.Core.Shopping_cart.DeleteShoppingCartProduct)
	api.Put("/shopping_cart/checkout", middleware.EmployeeAuthMiddleware(), handler.Core.Shopping_cart.CheckoutShoppingCart)
}
