package delivery

import (
	"strconv"
	"synapsis/config"
	"synapsis/internal/wrapper/usecase"
	"synapsis/pkg/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductsHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewProductsHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) ProductsHandler {
	return ProductsHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}

func (h ProductsHandler) GetAllProductByCategoryID(c *fiber.Ctx) error {

	init := exception.InitException(c, h.Conf, h.Log)

	categoryID, err := strconv.Atoi(c.Params("category_id"))

	if err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, "Invalid category_id", "category_id tidak valid", nil)
	}

	respData, errData := h.Usecase.Core.Products.GetAllProductByCategoryID(c.Context(), categoryID)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success get all product by category_id", "Sukses mengambil semua produk berdasarkan category_id", respData)
}
