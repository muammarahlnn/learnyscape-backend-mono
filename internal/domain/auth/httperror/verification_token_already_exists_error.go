package httperror

import (
	"errors"
	"learnyscape-backend-mono/internal/domain/auth/constant"
	"learnyscape-backend-mono/pkg/httperror"
	"net/http"
)

func NewVerificationTokenAlreadyExistsError() *httperror.ResponseError {
	msg := constant.VerificationTokenAlreadyExistsErrorMessage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusUnauthorized, msg)
}
