package httperror

import (
	"errors"
	"learnyscape-backend-mono/internal/domain/auth/constant"
	"learnyscape-backend-mono/pkg/httperror"
	"net/http"
)

func NewResetPasswordError() *httperror.ResponseError {
	msg := constant.ResetPasswordErrorMessage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusUnauthorized, msg)
}
