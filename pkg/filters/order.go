package filters

import (
	"fmt"

	"gorm.io/gorm"
)

type OrderType string

var (
	OrderAsc  = OrderType("asc")
	OrderDesc = OrderType("desc")
)

type Ordering struct {
	OrderBy string    `form:"order_by" json:"order_by"`
	Order   OrderType `form:"order" json:"order"`
}

func (p *Ordering) GetFilters() Filter {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s %s", p.OrderBy, p.Order))
	}
}
