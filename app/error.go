package app

import (
	"github.com/valyala/fasthttp"
)

type ExampleError struct {
	msg  string
	code int
}

func (c *ExampleError) Msg() string {
	return c.msg
}

func (c *ExampleError) Code() int {
	return c.code
}

func NewExampleError() *ExampleError {
	return &ExampleError{
		msg:  "error msg",
		code: fasthttp.StatusInternalServerError,
	}
}
