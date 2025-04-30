package httperror

import (
	"errors"
	"learnyscape-backend-mono/pkg/constant"
	"net/http"
)

func NewCanceledError() *ResponseError {
	msg := constant.RequestCanceledErrorMessage
	err := errors.New(msg)

	return NewResponseError(err, http.StatusRequestTimeout, msg)
}
