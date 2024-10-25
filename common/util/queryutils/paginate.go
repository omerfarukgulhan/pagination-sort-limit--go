package queryutils

import (
	"gorm.io/gorm"
	"math"
)

type Pagination struct {
	Limit      int         `json:"limit,omitempty" query:"limit"`
	Page       int         `json:"page,omitempty" query:"page"`
	Sort       string      `json:"sort,omitempty" query:"sort"`
	TotalRows  int64       `json:"totalRows"`
	TotalPages int         `json:"totalPages"`
	Rows       interface{} `json:"rows"`
}

func (pagination *Pagination) GetOffset() int {
	return (pagination.GetPage() - 1) * pagination.GetLimit()
}

func (pagination *Pagination) GetLimit() int {
	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	return pagination.Limit
}

func (pagination *Pagination) GetPage() int {
	if pagination.Page == 0 {
		pagination.Page = 1
	}
	return pagination.Page
}

func (pagination *Pagination) GetSort() string {
	if pagination.Sort == "" {
		pagination.Sort = "id desc"
	}
	return pagination.Sort
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
