package customers

import (
	"synapsis/internal/wrapper/handler"
	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Post("/customer/register", handler.Core.Customers.Register)
	api.Get("/customer/login", handler.Core.Customers.Login)
}