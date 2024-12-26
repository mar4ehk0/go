package order

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mar4ehk0/go/hw16_docker/internal/product"
	"github.com/mar4ehk0/go/hw16_docker/internal/user"
	"github.com/mar4ehk0/go/hw16_docker/pkg/db"
)

type RepoOrder struct {
	db *db.Connect
}

func NewRepoOrder(connect *db.Connect) *RepoOrder {
	return &RepoOrder{db: connect}
}

func (r *RepoOrder) AddWithTx(
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
		return 0, fmt.Errorf("insert order: %w", err)
	}

	for _, v := range products {
		_, err = tx.Exec("INSERT INTO orders_products (order_id, product_id) VALUES ($1, $2)", orderID, v.ID)
		if err != nil {
			return 0, fmt.Errorf("insert orders_products with order: %w", err)
		}
	}

	return orderID, nil
}

func (r *RepoOrder) GetByIDWithTx(tx *sqlx.Tx, id int) (Order, error) {
	var order Order
	err := tx.QueryRowx("SELECT id, user_id, order_date, total_amount FROM orders WHERE id=$1", id).StructScan(&order)
	if err != nil {
		return order, fmt.Errorf("select order id {%d}: %w", id, err)
	}

	if id != order.ID {
		return order, db.ErrDBNotFound
	}

	return order, nil
}

func (r *RepoOrder) DeleteProductsByIDWithTx(tx *sqlx.Tx, orderID int) error {
	res, err := tx.Exec("DELETE FROM orders_products WHERE order_id=$1", orderID)
	if err != nil {
		return fmt.Errorf("delete orders_products: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected orders_products for order {%d}: %w", orderID, err)
	}

	if count == 0 {
		return fmt.Errorf("was not delete orders_products related with order {%d}", orderID)
	}

	return nil
}

func (r *RepoOrder) AddProductsByIDWithTx(tx *sqlx.Tx, orderID int, productsID []int) error {
	for _, productID := range productsID {
		_, err := tx.Exec("INSERT INTO orders_products (order_id, product_id) VALUES ($1, $2)", orderID, productID)
		if err != nil {
			return fmt.Errorf("insert orders_products {%v} for order {%d}: %w", productsID, orderID, err)
		}
	}

	return nil
}

func (r *RepoOrder) UpdateTotalAmountOrderWithTX(tx *sqlx.Tx, orderID int, totalAmount int) error {
	_, err := tx.Exec("UPDATE orders SET total_amount=$1 WHERE id=$2", totalAmount, orderID)
	if err != nil {
		return fmt.Errorf("updated order {%d}: %w", orderID, err)
	}
	return nil
}
