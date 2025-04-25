package httperror

import "errors"

type ResponseError struct {
	err  error
	code int
	msg  string
}

func NewResponseError(err error, code int, msg string) *ResponseError {
	return &ResponseError{
		err:  err,
		code: code,
		msg:  msg,
	}
}

func (e ResponseError) Error() string {
	if e.msg == "" {
		return e.OriginalMessage()
	}

	return e.msg
}

func (e ResponseError) OriginalError() error {
	var currErr ResponseError
	currErr = e

	for {
		nextErr := currErr.err
		if nextErr == nil {
			break
		}

		var appErr ResponseError
		if !errors.As(nextErr, &appErr) {
			return nextErr
		}

		currErr = appErr
	}

	return e
}

func (e ResponseError) Code() int {
	return e.code
}

func (e ResponseError) OriginalMessage() string {
	return e.OriginalError().Error()
}

func (e ResponseError) Message() string {
	return e.msg
}
