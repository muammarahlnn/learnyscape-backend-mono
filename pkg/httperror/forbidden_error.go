package httperror

import (
	"errors"
	"learnyscape-backend-mono/pkg/constant"
	"net/http"
)

func NewForbiddenError() *ResponseError {
	msg := constant.ForbiddenErrorMessage
	err := errors.New(msg)

	return NewResponseError(err, http.StatusForbidden, msg)
}
