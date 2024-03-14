package core

import (
	"synapsis/config"
	"synapsis/internal/wrapper/repository"
	"synapsis/pkg/infra/db"

	"github.com/sirupsen/logrus"
	category	"synapsis/internal/core/category/usecase"
)

type CoreUsecase struct {
	Category	category.Usecase
}

func NewCoreUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreUsecase {
	return CoreUsecase{
		Category:	category.NewCategoryUsecase(repo, conf, dbList, log),
	}
}
