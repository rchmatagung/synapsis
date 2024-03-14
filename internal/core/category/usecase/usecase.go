package usecase

import (
	"context"
	"synapsis/config"
	"synapsis/internal/core/category/models"
	repo "synapsis/internal/wrapper/repository"
	"synapsis/pkg/exception"
	"synapsis/pkg/infra/db"

	"github.com/sirupsen/logrus"
)

type Usecase interface {
	GetAllCategory(context.Context) (*[]models.Category, *exception.Error)
}

type CategoryUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func NewCategoryUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) CategoryUsecase {
	return CategoryUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}

func (c CategoryUsecase) GetAllCategory(ctx context.Context) (*[]models.Category, *exception.Error) {

	response, err := c.Repo.Core.Category.GetAllCategory(ctx)
	if err != nil {
		return nil, exception.NewError(err.Code, err.Message, err.MessageInd)
	}

	return response, nil

}
