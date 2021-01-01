package http

import (
	"github.com/gin-gonic/gin"
	"home/internal/param"
	"home/internal/service"
	"home/pkg/context"
)

type Test struct {
}

func initTest(r *gin.RouterGroup) {
	test := r.Group("/test")
	{
		this := &Test{}
		test.GET("/SayHello", do(this.SayHello, false))
		test.GET("/UserInfo", do(this.UserInfo, false))
		test.GET("/UserList", do(this.UserList, false))
	}
}

// SayHello someone say hello
func (t *Test) SayHello(c *context.Context) {
	var p param.SayTest
	c.BindParam(&p)
	str, err := service.Test.SayHello(p)
	c.Response(str, err)
	//c.Response(service.Test.SayHello(p))
}

// UserInfo get user info
func (t *Test) UserInfo(c *context.Context) {
	c.Success(service.Test.UserInfo())
}

// UserList get user list
func (t *Test) UserList(c *context.Context) {
	c.Success(service.Test.UserList())
}
