package provider

import (
	"learnyscape-backend-mono/internal/system/handler"

	"github.com/gin-gonic/gin"
)

func BootstrapPing(router *gin.RouterGroup) {
	pingHandler := handler.NewPingHandler()

	pingHandler.Route(router)
}
