package queryutils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

type Filter struct {
	Field string
	Value interface{}
	Op    string
}

func ApplyFilter(filters []Filter) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, filter := range filters {
			switch filter.Op {
			case "=", "eq":
				db = db.Where(filter.Field+" = ?", filter.Value)
			case ">", "gt":
				db = db.Where(filter.Field+" > ?", filter.Value)
			case "<", "lt":
				db = db.Where(filter.Field+" < ?", filter.Value)
			case ">=", "gte":
				db = db.Where(filter.Field+" >= ?", filter.Value)
			case "<=", "lte":
				db = db.Where(filter.Field+" <= ?", filter.Value)
			case "!=", "ne":
				db = db.Where(filter.Field+" != ?", filter.Value)
			case "LIKE", "like":
				db = db.Where(filter.Field+" LIKE ?", "%"+filter.Value.(string)+"%")
			case "IN", "in":
				values := strings.Split(filter.Value.(string), ",")
				db = db.Where(filter.Field+" IN ?", values)
			case "NOT IN", "not_in":
				values := strings.Split(filter.Value.(string), ",")
				db = db.Where(filter.Field+" NOT IN ?", values)
			}
		}
		return db
	}
}

func ParseFilters(c *gin.Context) []Filter {
	filterQuery := c.Request.URL.Query()
	var filters []Filter
	for key, values := range filterQuery {
		if !strings.HasPrefix(key, "filter[") {
			continue
		}
		keyParts := strings.Split(key, "[")
		if len(keyParts) < 3 {
			continue
		}
		field := strings.TrimSuffix(keyParts[1], "]")
		op := strings.TrimSuffix(keyParts[2], "]")
		if len(values) > 0 {
			filters = append(filters, Filter{
				Field: field,
				Value: values[0],
				Op:    op,
			})
		}
	}
	return filters
}
