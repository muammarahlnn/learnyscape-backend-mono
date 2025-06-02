package pageutil

import (
	"learnyscape-backend-mono/pkg/dto"
	"math"
)

func NewMetadata(count, limit, page int64) *dto.PageMetaData {
	totalItems := count
	totalPage := int64(math.Ceil(float64(totalItems) / float64(limit)))

	return &dto.PageMetaData{
		Page:      page,
		Size:      limit,
		TotalItem: totalItems,
		TotalPage: totalPage,
	}
}
