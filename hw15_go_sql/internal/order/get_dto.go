package order

import (
	"time"

	"github.com/mar4ehk0/go/hw15_go_sql/internal/user"
)

type GetDto struct {
	ID          int       `json:"id"`
	User        user.User `json:"user"`
	OrderDate   time.Time `json: "order_date"`   //nolint:all
	TotalAmount int       `json: "total_amount"` //nolint:all
}
