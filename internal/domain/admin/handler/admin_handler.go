package handler

import (
	"learnyscape-backend-mono/internal/domain/admin/service"
	ginutil "learnyscape-backend-mono/pkg/util/gin"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	roleService service.AdminService
}

func NewAdminHandler(roleService service.AdminService) *AdminHandler {
	return &AdminHandler{
		roleService: roleService,
	}
}

func (h *AdminHandler) Route(router *gin.RouterGroup) {
	g := router.Group("/admin")
	{
		g.GET("/roles", h.GetAll)
	}
}

func (h *AdminHandler) GetAll(ctx *gin.Context) {
	res, err := h.roleService.GetAll(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseOK(ctx, res)
}
