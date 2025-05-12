package provider

import (
	"learnyscape-backend-mono/internal/domain/auth/handler"
	"learnyscape-backend-mono/internal/domain/auth/repository"
	"learnyscape-backend-mono/internal/domain/auth/service"

	"github.com/gin-gonic/gin"
)

func BootstrapAuth(router *gin.RouterGroup) {
	authDataStore := repository.NewAuthDataStore(dataStore)
	authService := service.NewAuthService(authDataStore, bcryptHasher, jwtUtil)
	authHandler := handler.NewAuthHandler(authService)

	authHandler.Route(router)
}
