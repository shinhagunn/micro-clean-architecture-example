package filters

import "gorm.io/gorm"

type Filter func(*gorm.DB) *gorm.DB

func Apply(db *gorm.DB, filters []Filter) *gorm.DB {
	for _, filter := range filters {
		filter(db)
	}

	return db
}
