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

func (r *Repo) AddWithTx(
	tx *sqlx.Tx,
	user user.User,
	products []product.Product,
	totalAmount int,
	date time.Time,
) (int, error) {
	var orderID int
	err := tx.QueryRow(
		"INSERT INTO orders (user_id, order_date, total_amount) VALUES ($1, $2, $3) RETURNING id",
		user.ID,
		date,
		totalAmount).Scan(&orderID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert order: %w", err)
	}

	for _, v := range products {
		_, err = tx.Exec("INSERT INTO orders_products (order_id, product_id) VALUES ($1, $2)", orderID, v.ID)
		if err != nil {
			return 0, fmt.Errorf("failed to insert order product: %w", err)
		}
	}

	return orderID, nil
}

func (r *Repo) GetByIDWithTx(tx *sqlx.Tx, id int) (Order, error) {
	var order Order
	err := tx.QueryRowx("SELECT id, user_id, order_date, total_amount FROM orders WHERE id=$1", id).StructScan(&order)
	if err != nil {
		return order, fmt.Errorf("failed to select order: %w", err)
	}

	return order, nil
}
