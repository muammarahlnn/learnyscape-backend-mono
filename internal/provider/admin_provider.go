package provider

import (
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/domain/admin/handler"
	"learnyscape-backend-mono/internal/domain/admin/repository"
	"learnyscape-backend-mono/internal/domain/admin/service"
	"learnyscape-backend-mono/pkg/constant"
	"learnyscape-backend-mono/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func BootstrapAdmin(cfg *config.Config, router *gin.RouterGroup) {
	adminDataStore := repository.NewAdminDataStore(dataStore)
	adminService := service.NewAdminService(
		cfg.Admin,
		adminDataStore,
		bcryptHasher,
		sendVerificationPublisher,
	)
	adminHandler := handler.NewAdminHandler(adminService)

	adminMiddleware := middleware.AuthMiddleware(jwtUtil, constant.AdminRole)
	adminHandler.Route(router, adminMiddleware)
}
