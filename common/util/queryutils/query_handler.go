package queryutils

import (
	"github.com/gin-gonic/gin"
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
