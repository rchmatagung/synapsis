package delivery

import (
	"fmt"
	"synapsis/config"
	"synapsis/internal/wrapper/usecase"
	"synapsis/pkg/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type NotFoundHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewNotFoundHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) NotFoundHandler {
	return NotFoundHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}

func (h NotFoundHandler) GetNotFound(c *fiber.Ctx) error {

	init := exception.InitException(c, h.Conf, h.Log)

	errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

	return exception.CreateResponse(init, fiber.StatusNotFound, errorMessage, errorMessage, nil)
}
