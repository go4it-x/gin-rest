package e

import "home/pkg/code"

type Error struct {
	Code uint32
	Msg  string
}

func (err Error) Error() string {
	if err.Msg == "" {
		err.Msg = code.Value(err.Code)
	}

	return err.Msg
}

func NewError(code uint32, msg string) Error {
	return Error{Code: code, Msg: msg}
}
