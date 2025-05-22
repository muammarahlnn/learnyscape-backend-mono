package httperror

import (
	"errors"
	"learnyscape-backend-mono/internal/domain/auth/constant"
	"learnyscape-backend-mono/pkg/httperror"
	"net/http"
)

func NewInvalidVerificationTokenError() *httperror.ResponseError {
	msg := constant.InvalidVerificationTokenErrorMessage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusUnauthorized, msg)
}
