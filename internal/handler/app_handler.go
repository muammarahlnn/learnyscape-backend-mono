package handler

import (
	"learnyscape-backend-mono/pkg/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (h *AppHandler) Route(r *gin.Engine) {
	r.NoRoute(h.routeNotFound)
	r.NoMethod(h.methodNotAllowed)
	r.GET("/ping", h.ping)
}

func (h *AppHandler) ping(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		dto.WebResponse[any]{
			Message: "pong",
		},
	)
}

func (h *AppHandler) routeNotFound(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		dto.WebResponse[any]{
			Message: "route not found",
		},
	)
}

func (h *AppHandler) methodNotAllowed(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		dto.WebResponse[any]{
			Message: "method not allowed",
		},
	)
}
