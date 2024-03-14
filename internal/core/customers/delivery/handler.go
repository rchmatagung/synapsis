package delivery

import (
	"synapsis/config"
	"synapsis/internal/core/customers/models"
	"synapsis/internal/wrapper/usecase"
	"synapsis/pkg/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CustomersHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewCustomersHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) CustomersHandler {
	return CustomersHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}

func (c CustomersHandler) Register(cf *fiber.Ctx) error {

	init := exception.InitException(cf, c.Conf, c.Log)

	dataReq := new(models.CustomerRegister)
	if err := cf.BodyParser(dataReq); err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, "Failed to parse request", "Gagal memparsing request", nil)
	}

	respData, errData := c.Usecase.Core.Customers.Register(cf.Context(), *dataReq)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusCreated, "Success register customer", "Sukses mendaftarkan customer", respData)

}

func (c CustomersHandler) Login(cf *fiber.Ctx) error {

	init := exception.InitException(cf, c.Conf, c.Log)

	username := cf.FormValue("username")
	password := cf.FormValue("password")

	respData, errData := c.Usecase.Core.Customers.Login(cf.Context(), username, password)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success login", "Sukses login", respData)

}
