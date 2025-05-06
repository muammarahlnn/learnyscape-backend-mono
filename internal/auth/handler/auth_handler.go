package handler

import (
	"learnyscape-backend-mono/internal/auth/dto"
	"learnyscape-backend-mono/internal/auth/service"
	ginutil "learnyscape-backend-mono/pkg/util/gin"

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

func (h *AuthHandler) Route(r *gin.RouterGroup) {
	g := r.Group("/auth")
	{
		g.POST("/login", h.login)
	}
}

func (h *AuthHandler) login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	res, err := h.authService.Login(ctx.Request.Context(), &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseOK(ctx, res)
}
