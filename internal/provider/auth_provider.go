package provider

import (
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/domain/auth/handler"
	"learnyscape-backend-mono/internal/domain/auth/mq"
	"learnyscape-backend-mono/internal/domain/auth/repository"
	"learnyscape-backend-mono/internal/domain/auth/service"

	"github.com/gin-gonic/gin"
)

func BootstrapAuth(cfg *config.Config, router *gin.RouterGroup) {
	accountVerifiedPublisher := mq.NewAccountVerifiedPublisher(rabbitmq)
	forgotPasswordPublisher := mq.NewForgotPasswordPublisher(rabbitmq)

	authDataStore := repository.NewAuthDataStore(dataStore)
	authService := service.NewAuthService(
		cfg.Auth,
		authDataStore,
		redisClient,
		bcryptHasher,
		jwtUtil,
		sendVerificationPublisher,
		accountVerifiedPublisher,
		forgotPasswordPublisher,
	)
	authHandler := handler.NewAuthHandler(authService)

	authHandler.Route(router)
}
