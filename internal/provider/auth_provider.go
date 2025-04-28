package provider

import (
	"learnyscape-backend-mono/internal/auth/handler"
	"learnyscape-backend-mono/internal/auth/repository"
	"learnyscape-backend-mono/internal/auth/service"

	"github.com/gin-gonic/gin"
)

func BootstrapAuth(router *gin.Engine) {
	authDataStore := repository.NewAuthDataStore(dataStore)
	authService := service.NewAuthService(authDataStore)
	authHandler := handler.NewAuthHandler(authService)

	authHandler.Route(router)
}
