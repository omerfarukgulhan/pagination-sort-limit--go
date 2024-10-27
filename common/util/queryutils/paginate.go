package queryutils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"strconv"
	"strings"
)

type Pagination struct {
	Limit      int         `json:"limit,omitempty" query:"limit"`
	Page       int         `json:"page,omitempty" query:"page"`
	Sort       string      `json:"sort,omitempty" query:"sort"`
	TotalRows  int64       `json:"totalRows"`
	TotalPages int         `json:"totalPages"`
	Data       interface{} `json:"data"`
}

func (pagination *Pagination) GetOffset() int {
	return (pagination.GetPage() - 1) * pagination.GetLimit()
}

func (pagination *Pagination) GetLimit() int {
	if pagination.Limit <= 0 {
		pagination.Limit = 10
	} else if pagination.Limit > 20 {
		pagination.Limit = 20
	}
	return pagination.Limit
}

func (pagination *Pagination) GetPage() int {
	if pagination.Page <= 0 {
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

func PaginateQuery(model interface{}, pagination *Pagination, db *gorm.DB) (func(db *gorm.DB) *gorm.DB, error) {
	if err := db.Model(model).Count(&pagination.TotalRows).Error; err != nil {
		return nil, err
	}

	pagination.TotalPages = int(math.Ceil(float64(pagination.TotalRows) / float64(pagination.GetLimit())))

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}, nil
}

func ParsePagination(c *gin.Context) (Pagination, error) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "20")
	sort := strings.ReplaceAll(c.DefaultQuery("sort", "id_desc"), "_", " ")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return Pagination{}, errors.New("invalid page parameter; must be a positive integer")
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		return Pagination{}, errors.New("invalid limit parameter; must be a positive integer")
	}
	pagination := Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
	return pagination, nil
}
