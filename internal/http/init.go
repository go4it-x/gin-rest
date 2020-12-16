package http

import (
	"github.com/douyu/jupiter/pkg/server/xgin"
	"github.com/gin-gonic/gin"
	"home/pkg/code"
	"home/pkg/code/e"
	"home/pkg/context"
	"home/pkg/middleware"
	"net/http"
)

// Init Load route
func Init(server *xgin.Server) {
	server.Use(middleware.Cors())
	server.NoRoute(do(func(c *context.Context) {
		c.Fail(e.Error{Code: code.NoRoute})
	}, false))
	server.GET("/", do(func(c *context.Context) {
		c.Success("hello world")
	}, false))

	v1 := server.Group("/v1")
	// init http router
	initTest(v1)
}

// do Custom processing handler function
func do(f func(*context.Context), isAuth bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer context.ErrorHandler(c)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		cc := new(context.Context)
		if isAuth {
			// TODO check jwt
			// c.Id = JwtId(ctx)
		}

		cc.Context = c
		f(cc)
	}
}
