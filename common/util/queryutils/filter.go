package queryutils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

type Filter struct {
	Field string
	Value interface{}
	Op    string
}

func FilterQuery(filters []Filter) (func(*gorm.DB) *gorm.DB, error) {
	var errs []error
	filterFunc := func(db *gorm.DB) *gorm.DB {
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
				valueStr, ok := filter.Value.(string)
				if !ok {
					errs = append(errs, fmt.Errorf("invalid value type for LIKE operation; expected string"))
					continue
				}
				db = db.Where(filter.Field+" LIKE ?", "%"+valueStr+"%")
			case "IN", "in":
				values, ok := filter.Value.(string)
				if !ok {
					errs = append(errs, fmt.Errorf("invalid value type for IN operation; expected comma-separated string"))
					continue
				}
				db = db.Where(filter.Field+" IN ?", strings.Split(values, ","))
			case "NOT IN", "not_in":
				values, ok := filter.Value.(string)
				if !ok {
					errs = append(errs, fmt.Errorf("invalid value type for NOT IN operation; expected comma-separated string"))
					continue
				}
				db = db.Where(filter.Field+" NOT IN ?", strings.Split(values, ","))
			default:
				errs = append(errs, fmt.Errorf("unsupported filter operation: %s", filter.Op))
			}
		}
		return db
	}
	if len(errs) > 0 {
		return filterFunc, fmt.Errorf("filter errors: %v", errs)
	}
	return filterFunc, nil
}

func ParseFilters(c *gin.Context) ([]Filter, error) {
	filterQuery := c.Request.URL.Query()
	var filters []Filter
	for key, values := range filterQuery {
		if !strings.HasPrefix(key, "filter[") {
			continue
		}
		keyParts := strings.Split(key, "[")
		if len(keyParts) < 3 {
			return nil, errors.New("invalid filter format; expected 'filter[field][op]'")
		}

		field := strings.TrimSuffix(keyParts[1], "]")
		op := strings.TrimSuffix(keyParts[2], "]")

		// Validate operation
		validOps := map[string]bool{"eq": true, "gt": true, "lt": true, "gte": true, "lte": true, "ne": true, "like": true, "in": true, "not_in": true}
		if _, ok := validOps[op]; !ok {
			return nil, errors.New("unsupported filter operation: " + op)
		}

		if len(values) > 0 {
			filters = append(filters, Filter{
				Field: field,
				Value: values[0],
				Op:    op,
			})
		}
	}
	return filters, nil
}
