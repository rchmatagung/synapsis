package delivery

import (
	"fmt"
	"strconv"
	"synapsis/config"
	"synapsis/internal/core/shopping_cart/models"
	"synapsis/internal/wrapper/usecase"
	"synapsis/pkg/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Shopping_cartHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewShopping_cartHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) Shopping_cartHandler {
	return Shopping_cartHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}

func (s Shopping_cartHandler) AddProductShoppingCart(c *fiber.Ctx) error {

	init := exception.InitException(c, s.Conf, s.Log)

	dateReq := new(models.ShoppingCartRequest)
	if err := c.BodyParser(dateReq); err != nil {
		return exception.CreateResponse_Log(init, fiber.StatusBadRequest, "Failed to parse request", "Gagal memparsing request", nil)
	}

	customerId := fmt.Sprintf("%v", c.Locals("customer_id"))
	customerIdInt, _ := strconv.ParseInt(customerId, 10, 64)

	username := fmt.Sprintf("%v", c.Locals("username"))

	respData, errData := s.Usecase.Core.Shopping_cart.AddProductShoppingCart(c.Context(), *dateReq, customerIdInt, username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusCreated, "Success add product to shopping cart", "Sukses menambahkan produk ke keranjang belanja", respData)

}

func (s Shopping_cartHandler) GetShoppingCart(c *fiber.Ctx) error {

	init := exception.InitException(c, s.Conf, s.Log)

	customerId := fmt.Sprintf("%v", c.Locals("customer_id"))
	customerIdInt, _ := strconv.ParseInt(customerId, 10, 64)

	respData, errData := s.Usecase.Core.Shopping_cart.GetShoppingCart(c.Context(), customerIdInt)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success get shopping cart", "Sukses mengambil keranjang belanja", respData)

}

func (s Shopping_cartHandler) DeleteShoppingCartProduct(c *fiber.Ctx) error {

	init := exception.InitException(c, s.Conf, s.Log)

	shoppingCartProductId := c.Params("shopping_cart_product_id")
	shoppingCartProductIdInt, _ := strconv.ParseInt(shoppingCartProductId, 10, 64)

	username := fmt.Sprintf("%v", c.Locals("username"))

	errData := s.Usecase.Core.Shopping_cart.DeleteShoppingCartProduct(c.Context(), shoppingCartProductIdInt, username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success delete product from shopping cart", "Sukses menghapus produk dari keranjang belanja", nil)

}

func (s Shopping_cartHandler) CheckoutShoppingCart(c *fiber.Ctx) error {

	init := exception.InitException(c, s.Conf, s.Log)

	customerId := fmt.Sprintf("%v", c.Locals("customer_id"))
	customerIdInt, _ := strconv.ParseInt(customerId, 10, 64)

	username := fmt.Sprintf("%v", c.Locals("username"))

	respData, errData := s.Usecase.Core.Shopping_cart.CheckoutShoppingCart(c.Context(), customerIdInt, username)
	if errData != nil {
		return exception.CreateResponse_Log(init, errData.Code, errData.Message, errData.MessageInd, nil)
	}

	return exception.CreateResponse_Log(init, fiber.StatusOK, "Success checkout shopping cart", "Sukses checkout keranjang belanja", respData)

}
