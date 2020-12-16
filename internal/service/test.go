package service

import (
	"fmt"
	"home/internal/model"
	"home/internal/param"
	"home/pkg/code"
	"home/pkg/code/e"
)

type TestService struct {
}

// NewTestService construct an instance
func NewTestService() *TestService {
	return &TestService{}
}

// SayHello say hello
func (t *TestService) SayHello(p param.SayTest) (str string, err e.Error) {
	// return err
	if p.Username == "admin" {
		err.Code = code.ParamsIsInvalid
		return
	}

	// Custom error msg
	if p.Username == "error" {
		err = e.NewError(44444444, "444444444 error")
		return
	}

	str = fmt.Sprintf("%s say hello world", p.Username)
	return
}

// UserInfo get user info
func (t *TestService) UserInfo() (user model.User) {
	user = model.User{
		Username: "admin",
		Password: "123456",
	}

	return
}

// UserList get user list
func (t *TestService) UserList() (users []model.User, count int) {
	users = []model.User{
		{Username: "test001", Password: "123456"},
		{Username: "test002", Password: "123456"},
		{Username: "test003", Password: "123456"},
	}

	count = len(users)
	return
}
