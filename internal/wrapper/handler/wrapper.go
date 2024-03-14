package handler

import (
	"synapsis/config"
	"synapsis/internal/wrapper/handler/core"
	"synapsis/internal/wrapper/handler/general"
	"synapsis/internal/wrapper/usecase"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	General general.GeneralHandler
	Core    core.CoreHandler
}

func NewHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) Handler {
	return Handler{
		General: general.NewGeneralHandler(uc, conf, log),
		Core:    core.NewCoreHandler(uc, conf, log),
	}

}
