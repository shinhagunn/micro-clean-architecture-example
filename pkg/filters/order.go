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
		orderBy := "created_at"
		order := OrderAsc

		if len(p.OrderBy) > 0 && p.OrderBy != orderBy {
			orderBy = p.OrderBy
		}

		if len(p.Order) > 0 && p.Order != order {
			order = p.Order
		}

		return db.Order(fmt.Sprintf("%s %s", orderBy, order))
	}
}
