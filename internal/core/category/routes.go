package category

import (
	"synapsis/internal/wrapper/handler"
	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
	api.Get("/category/getall", handler.Core.Category.GetAllCategory)
}