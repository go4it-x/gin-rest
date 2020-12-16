package engine

import (
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/server/xgin"
	"github.com/douyu/jupiter/pkg/util/xgo"
	"github.com/douyu/jupiter/pkg/xlog"
	"home/internal/http"
	"home/internal/service"
)

type Engine struct {
	jupiter.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		service.Init,
		xgo.ParallelWithError(
			eng.serveHTTP,
		),
	); err != nil {
		xlog.Panic("startup engine", xlog.Any("err", err))
	}

	return eng
}

func (eng *Engine) serveHTTP() error {
	server := xgin.StdConfig("http").Build()
	http.Init(server)
	return eng.Serve(server)
}
