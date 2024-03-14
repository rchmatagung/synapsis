package products

import (
	"synapsis/internal/middleware"
	"synapsis/internal/wrapper/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Get("/products/:category_id", middleware.EmployeeAuthMiddleware(), handler.Core.Products.GetAllProductByCategoryID)
}
