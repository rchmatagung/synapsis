package delivery

import (
	"synapsis/config"
	"synapsis/internal/wrapper/usecase"
	"synapsis/pkg/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CategoryHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewCategoryHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) CategoryHandler {
	return CategoryHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}

func (c CategoryHandler) GetAllCategory(cc *fiber.Ctx) error {

	init := exception.InitException(cc, c.Conf, c.Log)

	respData, errData := c.Usecase.Core.Category.GetAllCategory(cc.Context())
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success get all category", "Sukses mengambil semua kategori", respData)

}
