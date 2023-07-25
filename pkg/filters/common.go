package filters

import (
	"fmt"

	"gorm.io/gorm"
)

func OrderByAsc(field string) Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s asc", field))
	}
}

func OrderByDesc(field string) Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s desc", field))
	}
}

func WithLimit(limit int) Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

func WithOffset(offset int) Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	}
}
