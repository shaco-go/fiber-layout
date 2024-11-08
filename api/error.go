package api

import (
	"errors"
	"github.com/duke-git/lancet/v2/xerror"
)

var (
	ErrBadRequest = NewErr(1000, "Bad Request")
)

// Err 业务错误
type Err struct {
	code int
	msg  string
}

func (e *Err) Error() string {
	return e.msg
}

func (e *Err) Code() int {
	return e.code
}

// NewErr 创建业务错误
func NewErr(code int, msg string) error {
	return xerror.Unwrap(&Err{
		code: code,
		msg:  msg,
	})
}

// NewErrMsg 创建业务错误,只附加错误信息
func NewErrMsg(msg string) error {
	var err *Err
	_ = errors.As(ErrBadRequest, &err)
	return NewErr(err.Code(), msg)
}
