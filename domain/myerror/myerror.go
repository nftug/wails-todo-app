package myerror

import (
	"fmt"
)

type ErrorCode int

const (
	InvalidArgError ErrorCode = iota
	NotFoundError
)

type MyError struct {
	code    ErrorCode
	content *errorContent
	inner   error
}

type errorContent struct {
	Field   string `json:"field"`
	Message string `json:"error"`
}

func (e *MyError) Error() string { return e.inner.Error() }

func (e *MyError) Code() ErrorCode { return e.code }

func (e *MyError) Content() *errorContent { return e.content }

func New(c ErrorCode, e error) error {
	return &MyError{
		code:  c,
		inner: e,
	}
}

func Errorf(c ErrorCode, e error, field string, format string, a ...any) error {
	return &MyError{
		code:    c,
		inner:   e,
		content: &errorContent{field, fmt.Sprintf(format, a...)},
	}
}

func NewInvalidArgError(field string, format string, a ...any) error {
	return &MyError{
		code:    InvalidArgError,
		content: &errorContent{field, fmt.Sprintf(format, a...)},
	}
}
