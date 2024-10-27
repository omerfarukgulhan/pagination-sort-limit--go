package queryutils

import (
	"github.com/gin-gonic/gin"
)

type QueryHandler struct {
	Pagination Pagination `json:"pagination"`
	Filters    []Filter   `json:"filters"`
}

func QueryParser(c *gin.Context) QueryHandler {
	pagination := ParsePagination(c)
	filters := ParseFilters(c)
	queryHandler := QueryHandler{
		Pagination: pagination,
		Filters:    filters,
	}
	return queryHandler
}
