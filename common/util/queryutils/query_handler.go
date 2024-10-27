package queryutils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QueryHandler struct {
	Pagination Pagination `json:"pagination"`
	Filters    []Filter   `json:"filters"`
}

func QueryParser(c *gin.Context) (QueryHandler, error) {
	pagination, err := ParsePagination(c)
	if err != nil {
		return QueryHandler{}, err
	}
	filters, err := ParseFilters(c)
	if err != nil {
		return QueryHandler{}, err
	}
	queryHandler := QueryHandler{
		Pagination: pagination,
		Filters:    filters,
	}
	return queryHandler, nil
}

func ApplyQuery(db *gorm.DB, filters []Filter, pagination *Pagination, model interface{}) (*gorm.DB, error) {
	filterScope, err := FilterQuery(filters)
	if err != nil {
		return nil, err
	}

	db = db.Scopes(filterScope)
	paginationScope, err := PaginateQuery(model, pagination, db)
	if err != nil {
		return nil, err
	}

	db = db.Scopes(paginationScope)

	return db, nil
}
