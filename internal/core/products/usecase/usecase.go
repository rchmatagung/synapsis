package usecase

import (
	"context"
	"synapsis/config"
	"synapsis/internal/core/products/models"
	repo "synapsis/internal/wrapper/repository"
	"synapsis/pkg/exception"
	"synapsis/pkg/infra/db"

	"github.com/sirupsen/logrus"
)

type Usecase interface {
	GetAllProductByCategoryID(ctx context.Context, categoryId int) (*[]models.Product, *exception.Error)
}

type ProductsUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewProductsUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) ProductsUsecase {
	return ProductsUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}

func (u ProductsUsecase) GetAllProductByCategoryID(ctx context.Context, categoryId int) (*[]models.Product, *exception.Error) {

	response, err := u.Repo.Core.Products.GetAllProductByCategoryID(ctx, categoryId)
	if err != nil {
		return nil, exception.NewError(err.Code, err.Message, err.MessageInd)
	}

	return response, nil

}
