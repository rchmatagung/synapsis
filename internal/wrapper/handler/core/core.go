package core

import (
	"synapsis/config"
	"synapsis/internal/wrapper/usecase"

	"github.com/sirupsen/logrus"
)

type CoreHandler struct {
}

func NewCoreHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) CoreHandler {
	return CoreHandler{
	}
}
