package frsUtils

import (
	"math"
)

// Pagination common pagination struct
type Pagination struct {
	Page      int64 `json:"page"`
	TotalPage int64 `json:"total_page"`
	TotalData int64 `json:"total_data"`
}

// GetPaginationData return Pagination with page, total page, and total data
// total data can get from SELECT COUNT(*) FROM table_name query
func GetPaginationData(page, limit, totalData int) *Pagination {
	totalPage := math.Ceil(float64(totalData) / float64(limit))

	return &Pagination{
		Page:      int64(page),
		TotalPage: int64(totalPage),
		TotalData: int64(totalData),
	}
}
