package usecase

import (
	"context"
	"synapsis/config"
	"synapsis/internal/core/customers/models"
	repo "synapsis/internal/wrapper/repository"
	"synapsis/pkg/exception"
	"synapsis/pkg/infra/db"
	"synapsis/pkg/utils"
	myvalidator "synapsis/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Usecase interface {
	Register(ctx context.Context, req models.CustomerRegister) (*models.Customer, *exception.Error)
	Login(ctx context.Context, username, password string) (*models.LoginResponse, *exception.Error)
}

type CustomersUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewCustomersUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) CustomersUsecase {
	return CustomersUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}

func (c CustomersUsecase) Register(ctx context.Context, req models.CustomerRegister) (*models.Customer, *exception.Error) {

	errMsg, errMsgInd := myvalidator.ValidateDataRequest(req)

	if errMsg != "" || errMsgInd != "" {
		return nil, exception.NewError(fiber.StatusBadRequest, errMsg, errMsgInd)
	}

	response, err := c.Repo.Core.Customers.Register(ctx, req)
	if err != nil {
		return nil, exception.NewError(fiber.StatusInternalServerError, err.Message, err.MessageInd)
	}

	return response, nil

}

func (c CustomersUsecase) Login(ctx context.Context, username, password string) (*models.LoginResponse, *exception.Error) {

	response, err := c.Repo.Core.Customers.Login(ctx, username, password)
	if err != nil {
		return nil, exception.NewError(fiber.StatusBadRequest, err.Message, err.MessageInd)
	}

	token, _ := utils.GenereateJWT(c.Conf, response.CustomerId, response.Username)

	respData := &models.LoginResponse{
		CustomerId: response.CustomerId,
		Username:   response.Username,
		Password:   response.Password,
		Email:      response.Email,
		Token:      token,
	}

	return respData, nil

}
