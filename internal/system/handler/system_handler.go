package handler

import (
	"learnyscape-backend-mono/pkg/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemHandler struct {
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}

func (h *SystemHandler) Route(r *gin.Engine) {
	r.NoRoute(h.routeNotFound)
	r.NoMethod(h.methodNotAllowed)
}

func (h *SystemHandler) routeNotFound(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		dto.WebResponse[any]{
			Message: "route not found",
		},
	)
}

func (h *SystemHandler) methodNotAllowed(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		dto.WebResponse[any]{
			Message: "method not allowed",
		},
	)
}
