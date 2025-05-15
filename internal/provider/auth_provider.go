package provider

import (
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/domain/auth/handler"
	"learnyscape-backend-mono/internal/domain/auth/repository"
	"learnyscape-backend-mono/internal/domain/auth/service"
	"time"

	"github.com/gin-gonic/gin"
)

func BootstrapAuth(cfg *config.Config, router *gin.RouterGroup) {
	authDataStore := repository.NewAuthDataStore(dataStore)
	authService := service.NewAuthService(
		authDataStore,
		bcryptHasher,
		jwtUtil,
		redisClient,
		time.Duration(cfg.Redis.RefreshTokenExpiration)*time.Minute,
	)
	authHandler := handler.NewAuthHandler(authService)

	authHandler.Route(router)
}
