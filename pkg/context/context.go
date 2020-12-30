package context

import (
	"fmt"
	"github.com/douyu/jupiter/pkg/util/xtime"
	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/gin-gonic/gin"
	"home/pkg/code"
	"home/pkg/code/e"
	"home/pkg/validate"
	"net/http"
)

// Context Define a context
type Context struct {
	*gin.Context
	Id uint64 // Bind an id
}

// Ret Uniform results
type Ret struct {
	Code       uint32      `json:"code"`
	Msg        string      `json:"msg"`
	ServerTime int64       `json:"serverTime"`
	Data       interface{} `json:"data"`
}

type xRet struct {
	Ret
	Count interface{} `json:"count"`
	Ext   interface{} `json:"ext"`
}

// BindParam Binding parameters
func (c *Context) BindParam(params interface{}) {
	var err error
	if c.Context.Request.Method == "GET" {
		err = c.Context.ShouldBindQuery(params)
	} else {
		err = c.Context.ShouldBindJSON(params)
	}

	if err != nil && err.Error() != "EOF" {
		panic(e.Error{Code: code.Fail, Msg: err.Error()})
	}

	if msg := validate.Do(params); msg != "" {
		panic(e.Error{Code: code.Fail, Msg: msg})
	}
}

// Fail return fail
func (c *Context) Fail(err e.Error) {
	ret := Ret{
		Code:       err.Code,
		Msg:        err.Error(),
		ServerTime: xtime.GetTimestampInMilli(),
		Data:       "",
	}

	c.AbortWithStatusJSON(http.StatusOK, ret)
}

// Success return success
func (c *Context) Success(data ...interface{}) {
	length := len(data)
	var ret interface{}
	switch length {
	case 0, 1:
		tmp := Ret{
			Code:       code.Ok,
			Msg:        code.Value(code.Ok),
			ServerTime: xtime.GetTimestampInMilli(),
			Data:       "",
		}

		if length == 1 {
			tmp.Data = data[0]
		}

		ret = tmp
	case 2, 3:
		tmp := xRet{}
		tmp.Code = code.Ok
		tmp.Msg = code.Value(code.Ok)
		tmp.ServerTime = xtime.GetTimestampInMilli()
		tmp.Data = data[0]
		tmp.Count = data[1]
		tmp.Ext = ""
		if length == 3 {
			tmp.Ext = data[2]
		}

		ret = tmp
	}

	c.AbortWithStatusJSON(http.StatusOK, ret)
}

// ErrorHandler error handling
func ErrorHandler(c *gin.Context) {
	if err := recover(); err != nil {
		var ret Ret
		// Handling custom errors: panic
		if err, ok := err.(e.Error); ok {
			ret = Ret{
				Code:       err.Code,
				Msg:        err.Error(),
				ServerTime: xtime.GetTimestampInMilli(),
				Data:       "",
			}
		} else {
			// xlog.Error(string(debug.Stack()))
			msg := fmt.Sprintf("Server unknown error：%v", err)
			xlog.Error(msg)
			ret = Ret{
				Code:       code.ServerError,
				Msg:        msg,
				ServerTime: xtime.GetTimestampInMilli(),
				Data:       "",
			}
		}

		c.AbortWithStatusJSON(http.StatusOK, ret)
	}
}
