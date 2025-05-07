package provider

import (
	"learnyscape-backend-mono/internal/admin/handler"
	"learnyscape-backend-mono/internal/admin/repository"
	"learnyscape-backend-mono/internal/admin/service"

	"github.com/gin-gonic/gin"
)

func BootstrapAdmin(router *gin.RouterGroup) {
	adminDataStore := repository.NewAdminDataStore(dataStore)
	adminService := service.NewAdminService(adminDataStore)
	adminHandler := handler.NewAdminHandler(adminService)

	adminHandler.Route(router)
}
