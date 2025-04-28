package provider

import (
	"learnyscape-backend-mono/internal/config"
	"learnyscape-backend-mono/internal/handler"

	"github.com/gin-gonic/gin"
)

func BootstrapHttp(cfg *config.Config, router *gin.Engine) {
	appHandler := handler.NewAppHandler()
	appHandler.Route(router)

	BootstrapAuth(router)
}
