package core

import (
	"synapsis/config"
	"synapsis/pkg/infra/db"

	"github.com/sirupsen/logrus"
	category	"synapsis/internal/core/category/repository"
	customers	"synapsis/internal/core/customers/repository"
)

type CoreRepository struct {
	Category	category.Repository
	Customers	customers.Repository
}

func NewCoreRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreRepository {
	return CoreRepository{
		Category:	category.NewCategoryRepo(dbList),
		Customers:	customers.NewCustomersRepo(dbList),
	}
}
