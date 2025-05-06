package handler

import (
	"learnyscape-backend-mono/pkg/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Route(r *gin.RouterGroup) {
	r.GET("/ping", h.ping)
}

func (h *PingHandler) ping(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		dto.WebResponse[any]{
			Message: "pong",
		},
	)
}
