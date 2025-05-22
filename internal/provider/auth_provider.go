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
	sendVerificationProducer := mq.NewSendVerificationPublisher(rabbitmq)
	accountVerifiedProducer := mq.NewAccountVerifiedPublisher(rabbitmq)

	authDataStore := repository.NewAuthDataStore(dataStore)
	authService := service.NewAuthService(
		authDataStore,
		bcryptHasher,
		jwtUtil,
		redisClient,
		cfg.Auth,
		sendVerificationProducer,
		accountVerifiedProducer,
	)
	authHandler := handler.NewAuthHandler(authService)

	authHandler.Route(router)
}
