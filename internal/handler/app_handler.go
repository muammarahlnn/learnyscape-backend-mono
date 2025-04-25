package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (h *AppHandler) Route(r *gin.Engine) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "gokss air verse is cooking here!")
	})
}
