package order

import "time"

type Order struct {
	ID          int       `json:"id"`
	UserID      int       `db:"user_id" json:"user_id"`            //nolint:all
	OrderDate   time.Time `db:"order_date" json: "order_date"`     //nolint:all
	TotalAmount int       `db:"total_amount" json: "total_amount"` //nolint:all
}
