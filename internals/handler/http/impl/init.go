package impl

import (
	"github.com/bagusandrian/framework/internals/config"
	httpframework "github.com/bagusandrian/framework/internals/handler/http"
	ucDummy "github.com/bagusandrian/framework/internals/usecase/dummy"
)

type handler struct {
	config  *config.Config
	usecase ucDummy.Usecase
}

func New(
	cfg *config.Config, ucDummy ucDummy.Usecase,
) httpframework.Handler {

	// init repository
	h := &handler{
		config:  cfg,
		usecase: ucDummy,
	}
	return h
}
