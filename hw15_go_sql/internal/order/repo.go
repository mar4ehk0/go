package order

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mar4ehk0/go/hw15_go_sql/internal/product"
	"github.com/mar4ehk0/go/hw15_go_sql/internal/user"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Add(tx *sqlx.Tx, user user.User, products []product.Product, totalAmount int, date time.Time) (int, error) {
	var orderId int
	err := tx.QueryRow(
		"INSERT INTO orders (user_id, order_date, total_amount) VALUES ($1, $2, $3) RETURNING id",
		user.Id,
		date,
		totalAmount).Scan(&orderId)
	if err != nil {
		return 0, fmt.Errorf("failed to insert order: %v", err)
	}

	for _, v := range products {
		_, err := tx.Exec("INSERT INTO orders_products (order_id, product_id) VALUES ($1, $2)", orderId, v.Id)
		if err != nil {
			return 0, fmt.Errorf("failed to insert order product: %v", err)
		}
	}

	return orderId, nil
}
