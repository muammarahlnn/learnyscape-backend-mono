package httperror

import (
	"errors"
	"learnyscape-backend-mono/internal/domain/admin/constant"
	"learnyscape-backend-mono/pkg/httperror"
	"net/http"
)

func NewUserAlreadyExistsError() *httperror.ResponseError {
	msg := constant.UserAlreadyExistsErrorMessage
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusConflict, msg)
}
