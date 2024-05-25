package interfaces

import (
	"fmt"
)

type ErrorCode int

const (
	InvalidArgError ErrorCode = iota
	NotFoundError
)

type ErrorResponse struct {
	code    ErrorCode
	content *errorContent
	inner   error
}

type errorContent struct {
	Field   string `json:"field"`
	Message string `json:"error"`
}

func (e *ErrorResponse) Error() string { return e.inner.Error() }

func (e *ErrorResponse) Code() ErrorCode { return e.code }

func (e *ErrorResponse) Content() *errorContent { return e.content }

func New(c ErrorCode, e error) error {
	return &ErrorResponse{
		code:  c,
		inner: e,
	}
}

func Errorf(c ErrorCode, e error, field string, format string, a ...any) error {
	return &ErrorResponse{
		code:    c,
		inner:   e,
		content: &errorContent{field, fmt.Sprintf(format, a...)},
	}
}

func NewInvalidArgError(field string, format string, a ...any) error {
	return &ErrorResponse{
		code:    InvalidArgError,
		content: &errorContent{field, fmt.Sprintf(format, a...)},
		inner:   fmt.Errorf(format, a...),
	}
}

func NewNotFoundError(itemName string) error {
	return &ErrorResponse{
		code:    NotFoundError,
		content: &errorContent{itemName, "Not found."},
		inner:   fmt.Errorf("%s is not found", itemName),
	}
}
