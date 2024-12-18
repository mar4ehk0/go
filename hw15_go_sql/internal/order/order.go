package order

import "time"

type Order struct {
	Id          int       `json:"id"`
	UserId      int       `db:"user_id" json:"user_id"`
	OrderDate   time.Time `db:"order_date" json: "order_date"`
	TotalAmount int       `db:"total_amount" json: "total_amount"`
}
