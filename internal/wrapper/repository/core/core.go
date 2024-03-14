package core

import (
	"synapsis/config"
	"synapsis/pkg/infra/db"

	"github.com/sirupsen/logrus"
	category	"synapsis/internal/core/category/repository"
	customers	"synapsis/internal/core/customers/repository"
	products	"synapsis/internal/core/products/repository"
	shopping_cart	"synapsis/internal/core/shopping_cart/repository"
)

type CoreRepository struct {
	Category	category.Repository
	Customers	customers.Repository
	Products	products.Repository
	Shopping_cart	shopping_cart.Repository
}

func NewCoreRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreRepository {
	return CoreRepository{
		Category:	category.NewCategoryRepo(dbList),
		Customers:	customers.NewCustomersRepo(dbList),
		Products:	products.NewProductsRepo(dbList),
		Shopping_cart:	shopping_cart.NewShopping_cartRepo(dbList),
	}
}
