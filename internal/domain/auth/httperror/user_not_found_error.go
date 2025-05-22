package httperror

import (
	"errors"
	"learnyscape-backend-mono/internal/domain/auth/constant"
	"learnyscape-backend-mono/pkg/httperror"
	"net/http"
)

func NewUserNotFoundError() *httperror.ResponseError {
	msg := constant.UserNotFoundErrorMessage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusNotFound, msg)
}
