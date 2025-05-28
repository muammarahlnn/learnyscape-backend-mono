package handler

import (
	"learnyscape-backend-mono/internal/domain/admin/dto"
	"learnyscape-backend-mono/internal/domain/admin/service"
	ginutil "learnyscape-backend-mono/pkg/util/gin"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService service.AdminService
}

func NewAdminHandler(adminService service.AdminService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

func (h *AdminHandler) Route(
	router *gin.RouterGroup,
	adminMiddleware gin.HandlerFunc,
) {
	g := router.Group("/admin", adminMiddleware)
	{
		g.GET("/roles", h.getAllRoles)
		g.POST("/users", h.createUser)
		g.GET("/users", h.getAllUsers)
	}
}

func (h *AdminHandler) getAllRoles(ctx *gin.Context) {
	res, err := h.adminService.GetRoles(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseOK(ctx, res)
}

func (h *AdminHandler) createUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	res, err := h.adminService.CreateUser(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseOK(ctx, res)
}

func (h *AdminHandler) getAllUsers(ctx *gin.Context) {
	res, err := h.adminService.GetAllUsers(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseOK(ctx, res)
}
