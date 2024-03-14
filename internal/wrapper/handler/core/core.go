package core

import (
	"synapsis/config"
	"synapsis/internal/wrapper/usecase"

	"github.com/sirupsen/logrus"
	category	"synapsis/internal/core/category/delivery"
)

type CoreHandler struct {
	Category	category.CategoryHandler
}

func NewCoreHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) CoreHandler {
	return CoreHandler{
		Category:	category.NewCategoryHandler(uc, conf, log),
	}
}
