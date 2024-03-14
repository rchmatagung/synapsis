package core

import (
	"synapsis/config"
	"synapsis/internal/wrapper/repository"
	"synapsis/pkg/infra/db"

	"github.com/sirupsen/logrus"
	category	"synapsis/internal/core/category/usecase"
	customers	"synapsis/internal/core/customers/usecase"
	products	"synapsis/internal/core/products/usecase"
	shopping_cart	"synapsis/internal/core/shopping_cart/usecase"
)

type CoreUsecase struct {
	Category	category.Usecase
	Customers	customers.Usecase
	Products	products.Usecase
	Shopping_cart	shopping_cart.Usecase
}

func NewCoreUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreUsecase {
	return CoreUsecase{
		Category:	category.NewCategoryUsecase(repo, conf, dbList, log),
		Customers:	customers.NewCustomersUsecase(repo, conf, dbList, log),
		Products:	products.NewProductsUsecase(repo, conf, dbList, log),
		Shopping_cart:	shopping_cart.NewShopping_cartUsecase(repo, conf, dbList, log),
	}
}
