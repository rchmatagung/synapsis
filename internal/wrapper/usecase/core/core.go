package core

import (
	"synapsis/config"
	"synapsis/internal/wrapper/repository"
	"synapsis/pkg/infra/db"

	"github.com/sirupsen/logrus"
	category	"synapsis/internal/core/category/usecase"
	customers	"synapsis/internal/core/customers/usecase"
)

type CoreUsecase struct {
	Category	category.Usecase
	Customers	customers.Usecase
}

func NewCoreUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreUsecase {
	return CoreUsecase{
		Category:	category.NewCategoryUsecase(repo, conf, dbList, log),
		Customers:	customers.NewCustomersUsecase(repo, conf, dbList, log),
	}
}
