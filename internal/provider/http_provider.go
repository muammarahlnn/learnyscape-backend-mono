package provider

import (
	"learnyscape-backend-mono/internal/config"

	"github.com/gin-gonic/gin"
)

func BootstrapHttp(cfg *config.Config, router *gin.Engine) {
	BootstrapSystem(router)

	rg := router.Group("/api/v1")
	BootstrapPing(rg)
	BootstrapAuth(rg)
}
