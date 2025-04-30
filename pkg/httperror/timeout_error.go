package httperror

import (
	"errors"
	"learnyscape-backend-mono/pkg/constant"
	"net/http"
)

func NewTimeoutError() *ResponseError {
	msg := constant.RequestTimeoutErrorMessage
	err := errors.New(msg)

	return NewResponseError(err, http.StatusGatewayTimeout, msg)
}
