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
		g.POST("/refresh", h.refresh)
		g.POST("/verify", h.verify)
		g.POST("/resend-verification", h.resendVerification)
		g.POST("/forgot-password", h.forgotPassword)
		g.PUT("/reset-password", h.resetPassword)
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

func (h *AuthHandler) resendVerification(ctx *gin.Context) {
	var req dto.ResendVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	if err := h.authService.ResendVerification(ctx, &req); err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseNoContent(ctx)
}

func (h *AuthHandler) forgotPassword(ctx *gin.Context) {
	var req dto.ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	if err := h.authService.ForgotPassword(ctx, &req); err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseNoContent(ctx)
}

func (h *AuthHandler) resetPassword(ctx *gin.Context) {
	var req dto.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	if err := h.authService.ResetPassword(ctx, &req); err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseNoContent(ctx)
}
