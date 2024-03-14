package core

import (
	"synapsis/config"
	"synapsis/internal/wrapper/usecase"

	"github.com/sirupsen/logrus"
	category	"synapsis/internal/core/category/delivery"
	customers	"synapsis/internal/core/customers/delivery"
)

type CoreHandler struct {
	Category	category.CategoryHandler
	Customers	customers.CustomersHandler
}

func NewCoreHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) CoreHandler {
	return CoreHandler{
		Category:	category.NewCategoryHandler(uc, conf, log),
		Customers:	customers.NewCustomersHandler(uc, conf, log),
	}
}
