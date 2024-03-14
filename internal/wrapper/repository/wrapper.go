package repository

import (
	"synapsis/config"
	"synapsis/internal/wrapper/repository/core"
	"synapsis/internal/wrapper/repository/general"
	"synapsis/pkg/infra/db"

	"github.com/sirupsen/logrus"
)

type Repository struct {
	General general.GeneralRepository
	Core    core.CoreRepository
}

func NewRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) Repository {
	return Repository{
		General: general.NewGeneralRepository(conf, dbList, log),
		Core:    core.NewCoreRepository(conf, dbList, log),
	}
}
