package general

import (
	"synapsis/config"
	notfound "synapsis/internal/general/notfound/delivery"
	root "synapsis/internal/general/root/delivery"
	"synapsis/internal/wrapper/usecase"

	"github.com/sirupsen/logrus"
)

type GeneralHandler struct {
	Root     root.RootHandler
	NotFound notfound.NotFoundHandler
}

func NewGeneralHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) GeneralHandler {
	return GeneralHandler{
		Root:     root.NewRootHandler(uc, conf, log),
		NotFound: notfound.NewNotFoundHandler(uc, conf, log),
	}
}
