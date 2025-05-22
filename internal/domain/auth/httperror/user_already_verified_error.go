package httperror

import (
	"errors"
	"learnyscape-backend-mono/internal/domain/auth/constant"
	"learnyscape-backend-mono/pkg/httperror"
	"net/http"
)

func NewUserAlreadyVerifiedError() *httperror.ResponseError {
	msg := constant.UserAlreadyVerifiedErrorMessage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusConflict, msg)
}
