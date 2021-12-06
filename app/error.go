package app

import "github.com/valyala/fasthttp"

type customError struct {
	Msg  string `json:"Error"`
	Code int `json:"-"`
}

func (c customError) ReturnedMsg() string {
	return c.Msg
}

func (c customError) ReturnedCode() int {
	return c.Code
}

func NewExampleError() *customError {
	return &customError{
		Msg:  "example error msg",
		Code: fasthttp.StatusInternalServerError,
	}
}

func NewExampleError2() *customError {
	return &customError{
		Msg:  "example error msg2",
		Code: fasthttp.StatusInternalServerError,
	}
}
