package handler

import (
	"learnyscape-backend-mono/internal/auth/service"
	"learnyscape-backend-mono/pkg/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Route(r *gin.Engine) {
	r.GET("/auth/test", h.test)
}

func (h *AuthHandler) test(ctx *gin.Context) {
	res, err := h.authService.Test(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.WebResponse[string]{Message: res})
}
