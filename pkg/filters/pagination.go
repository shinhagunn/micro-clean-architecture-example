package filters

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Page  int   `form:"page" json:"page"`
	Limit int   `form:"limit" json:"limit"`
	Total int64 `form:"-" json:"total"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetFilters() Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetLimit())
	}
}
