package provider

import (
	"learnyscape-backend-mono/internal/system/handler"

	"github.com/gin-gonic/gin"
)

func BootstrapSystem(router *gin.Engine) {
	systemHandler := handler.NewSystemHandler()

	systemHandler.Route(router)
}
