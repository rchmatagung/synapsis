package general

import (
	"synapsis/config"
	"synapsis/internal/wrapper/repository"
	"synapsis/pkg/infra/db"

	"github.com/sirupsen/logrus"
)

type GeneralUsecase struct {
}

func NewGeneralUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) GeneralUsecase {
	return GeneralUsecase{}
}
