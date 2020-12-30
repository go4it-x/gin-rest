package engine

import (
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/server/xgin"
	"github.com/douyu/jupiter/pkg/store/gorm"
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
		eng.serveDB,
		xgo.ParallelWithError(
			eng.serveHTTP,
		),
	); err != nil {
		xlog.Panic("startup engine", xlog.Any("err", err))
	}

	return eng
}

func (eng *Engine) serveDB() error {
	db := gorm.StdConfig("test").Build()
	service.Init(db)
	return nil
}

func (eng *Engine) serveHTTP() error {
	server := xgin.StdConfig("http").Build()
	http.Init(server)
	return eng.Serve(server)
}
