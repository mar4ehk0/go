package order

import "time"

type Order struct {
	Id          int       `json:"id"`                              //nolint:structtag
	UserId      int       `db:"user_id" json:"user_id"`            //nolint:structtag
	OrderDate   time.Time `db:"order_date" json: "order_date"`     //nolint:structtag
	TotalAmount int       `db:"total_amount" json: "total_amount"` //nolint:structtag
}
