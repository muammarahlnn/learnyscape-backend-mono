package handler

import (
	"learnyscape-backend-mono/internal/domain/auth/dto"
	"learnyscape-backend-mono/internal/domain/auth/service"
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
		g.POST("/register", h.register)
		g.POST("/refresh", h.refresh)
		g.POST("/verify", h.verify)
	}
}

func (h *AuthHandler) login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	res, err := h.authService.Login(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseOK(ctx, res)
}

func (h *AuthHandler) register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	res, err := h.authService.Register(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseOK(ctx, res)
}

func (h *AuthHandler) refresh(ctx *gin.Context) {
	var req dto.RefreshRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	res, err := h.authService.Refresh(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseOK(ctx, res)
}

func (h *AuthHandler) verify(ctx *gin.Context) {
	var req dto.VerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	res, err := h.authService.Verify(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseOK(ctx, res)
}
