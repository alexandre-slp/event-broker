package infra

import (
	"context"
	"runtime/debug"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type customError struct {
	msg        string
	statusCode codes.Code
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
	ce, ok := p.(*customError)
	if ok {
		logger.Error().Msg(ce.msg)
		return status.Errorf(codes.Internal, "error: %v", ce.msg)
	}

	logger.Error().Msg(string(debug.Stack()))
	return status.Errorf(codes.Internal, "unhadled error: %v", p)
}

//NewExampleError Example error
func NewExampleError() *customError {
	errorMsg := "example error msg"
	return &customError{
		msg:        errorMsg,
		statusCode: codes.Internal,
	}
}
