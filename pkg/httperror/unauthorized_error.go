package httperror

import (
	"errors"
	"learnyscape-backend-mono/pkg/constant"
	"net/http"
)

func NewUnauthorizedError() *ResponseError {
	msg := constant.UnauthorizedErrorMessage
	err := errors.New(msg)

	return NewResponseError(err, http.StatusUnauthorized, msg)
}
