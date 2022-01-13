package app

import (
	"context"
	"encoding/json"
	"github.com/rotisserie/eris"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type customError struct {
	msg        string
	statusCode codes.Code
	erisError  error
}

func (ce customError) Error() string {
	return ce.msg
}

func (ce customError) returnedCode() codes.Code {
	return ce.statusCode
}

// CustomPanicHandler Define how to handle panic
func CustomPanicHandler(ctx context.Context, p interface{}) (err error) {
	logger := ctx.Value("logger").(zerolog.Logger)
	jsonFormatted, err := json.MarshalIndent(eris.ToJSON(p.(error), true), "", "  ")
	if err != nil {
		return err
	}
	logger.Error().Msg(string(jsonFormatted))
	return status.Errorf(codes.Internal, "panic triggered: %v", p)
}

//NewExampleError Example error
func NewExampleError(err error) *customError {
	errorMsg := "example error msg"
	return &customError{
		msg:        errorMsg,
		statusCode: codes.Internal,
		erisError:  err,
	}
}
