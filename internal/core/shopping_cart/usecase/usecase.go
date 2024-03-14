package usecase

import (
	"context"
	"synapsis/config"
	"synapsis/internal/core/shopping_cart/models"
	repo "synapsis/internal/wrapper/repository"
	"synapsis/pkg/exception"
	"synapsis/pkg/infra/db"

	"github.com/sirupsen/logrus"
)

type Usecase interface {
	AddProductShoppingCart(ctx context.Context, req models.ShoppingCartRequest, customerId int64, createdBy string) (*models.ShoppingCart, *exception.Error)
	GetShoppingCart(ctx context.Context, customerId int64) (*models.ShoppingCart, *exception.Error)
	DeleteShoppingCartProduct(ctx context.Context, shoppingCartProductId int64, deletedBy string) *exception.Error
	CheckoutShoppingCart(ctx context.Context, customerId int64, updatedBy string) (*models.ShoppingCart, *exception.Error)
}

type Shopping_cartUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewShopping_cartUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) Shopping_cartUsecase {
	return Shopping_cartUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}

func (s Shopping_cartUsecase) AddProductShoppingCart(ctx context.Context, req models.ShoppingCartRequest, customerId int64, createdBy string) (*models.ShoppingCart, *exception.Error) {

	shoppingCart, err := s.Repo.Core.Shopping_cart.AddProductShoppingCart(ctx, req, customerId, createdBy)
	if err != nil {
		return nil, exception.NewError(err.Code, err.Message, err.MessageInd)
	}

	return shoppingCart, nil

}

func (s Shopping_cartUsecase) GetShoppingCart(ctx context.Context, customerId int64) (*models.ShoppingCart, *exception.Error) {

	shoppingCart, err := s.Repo.Core.Shopping_cart.GetShoppingCart(ctx, customerId)
	if err != nil {
		return nil, exception.NewError(err.Code, err.Message, err.MessageInd)
	}

	return shoppingCart, nil

}

func (s Shopping_cartUsecase) DeleteShoppingCartProduct(ctx context.Context, shoppingCartProductId int64, deletedBy string) *exception.Error {

	err := s.Repo.Core.Shopping_cart.DeleteShoppingCartProduct(ctx, shoppingCartProductId, deletedBy)
	if err != nil {
		return exception.NewError(err.Code, err.Message, err.MessageInd)
	}

	return nil

}

func (s Shopping_cartUsecase) CheckoutShoppingCart(ctx context.Context, customerId int64, updatedBy string) (*models.ShoppingCart, *exception.Error) {

	shoppingCart, err := s.Repo.Core.Shopping_cart.CheckoutShoppingCart(ctx, customerId, updatedBy)
	if err != nil {
		return nil, exception.NewError(err.Code, err.Message, err.MessageInd)
	}

	return shoppingCart, nil

}