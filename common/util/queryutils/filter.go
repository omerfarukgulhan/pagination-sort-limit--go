package queryutils

import (
	"gorm.io/gorm"
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
			case "=":
				db = db.Where(filter.Field+" = ?", filter.Value)
			case ">":
				db = db.Where(filter.Field+" > ?", filter.Value)
			case "<":
				db = db.Where(filter.Field+" < ?", filter.Value)
			case ">=":
				db = db.Where(filter.Field+" >= ?", filter.Value)
			case "<=":
				db = db.Where(filter.Field+" <= ?", filter.Value)
			case "!=":
				db = db.Where(filter.Field+" != ?", filter.Value)
			case "LIKE":
				db = db.Where(filter.Field+" LIKE ?", filter.Value)
			case "IN":
				db = db.Where(filter.Field+" IN ?", filter.Value)
			case "NOT IN":
				db = db.Where(filter.Field+" NOT IN ?", filter.Value)
			default:
			}
		}
		return db
	}
}
