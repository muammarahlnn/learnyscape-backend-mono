package handler

import (
	"fmt"
	"learnyscape-backend-mono/internal/domain/admin/dto"
	"learnyscape-backend-mono/internal/domain/admin/service"
	ginutil "learnyscape-backend-mono/pkg/util/gin"
	pageutil "learnyscape-backend-mono/pkg/util/page"
	"reflect"

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
		g.GET("/users", h.search)
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

func (h *AdminHandler) search(ctx *gin.Context) {
	pagination := ginutil.ParsePagination(ctx)
	req := &dto.SearchUserRequest{Pagination: pagination}
	if err := ctx.ShouldBindQuery(req); err != nil {
		fmt.Println("masukk:", reflect.TypeOf(err))
		ctx.Error(err)
		return
	}

	res, paging, err := h.adminService.SearchUser(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}
	paging.Links = pageutil.NewLinks(
		ctx.Request,
		int(paging.Page),
		int(paging.Size),
		int(paging.TotalItem),
		int(paging.TotalPage),
	)

	ginutil.ResponsePagination(ctx, res, paging)
}
