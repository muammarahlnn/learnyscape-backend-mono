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
		g.GET("/users", h.searchUser)
		g.GET("/users/:id", h.getUser)
		g.PUT("/users/:id", h.updateUser)
		g.DELETE("/users/:id", h.deleteUser)
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

func (h *AdminHandler) searchUser(ctx *gin.Context) {
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

func (h *AdminHandler) getUser(ctx *gin.Context) {
	var pathParams dto.GetUserPathParams
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.Error(err)
		return
	}

	res, err := h.adminService.GetUser(ctx, pathParams.ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseOK(ctx, res)
}

func (h *AdminHandler) updateUser(ctx *gin.Context) {
	var pathParams dto.UpdateUserPathParams
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.Error(err)
		return
	}

	var req dto.UpdaetUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	res, err := h.adminService.UpdateUser(ctx, pathParams.ID, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseOK(ctx, res)
}

func (h *AdminHandler) deleteUser(ctx *gin.Context) {
	var pathParams dto.DeleteUserPathParams
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.Error(err)
		return
	}

	err := h.adminService.DeleteUser(ctx, pathParams.ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutil.ResponseNoContent(ctx)
}
