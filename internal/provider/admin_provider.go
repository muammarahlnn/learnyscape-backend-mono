package provider

import (
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/domain/admin/handler"
	"learnyscape-backend-mono/internal/domain/admin/repository"
	"learnyscape-backend-mono/internal/domain/admin/service"

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

	adminHandler.Route(router)
}
