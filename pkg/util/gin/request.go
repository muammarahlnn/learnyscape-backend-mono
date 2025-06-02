package ginutil

import (
	"learnyscape-backend-mono/pkg/constant"
	"learnyscape-backend-mono/pkg/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParsePagination(ctx *gin.Context) *dto.Pagination {
	limitStr := ctx.DefaultQuery("limit", strconv.Itoa(constant.DefaultLimit))
	pageStr := ctx.DefaultQuery("page", strconv.Itoa(constant.DefaultPage))

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit <= 0 {
		limit = constant.DefaultLimit
	}

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil || page <= 0 {
		page = constant.DefaultPage
	}

	return &dto.Pagination{
		Limit: limit,
		Page:  page,
	}
}
