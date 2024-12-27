package product

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mar4ehk0/go/hw16_docker/pkg/db"
)

type RepoProduct struct {
	db *db.Connect
}

func NewRepoProduct(connect *db.Connect) *RepoProduct {
	return &RepoProduct{db: connect}
}

func (r *RepoProduct) Add(dto *EntryDto) (int, error) {
	query := "INSERT INTO products (name, price) VALUES (:name, :price) RETURNING id"
	stmt, err := r.db.Connect.PrepareNamed(query)
	if err != nil {
		wrappedErr := fmt.Errorf("prepare query insert product {%v}: %w", dto, err)
		return 0, wrappedErr
	}

	var id int
	err = stmt.Get(&id, dto)
	if err != nil {
		if r.db.IsErrDuplicate(err) {
			return 0, fmt.Errorf("already exist product{%v}: %w", dto, errors.Join(db.ErrDBDuplicateKey, err))
		}
		return 0, fmt.Errorf("sql insert product{%v}: %w", dto, err)
	}

	return id, nil
}

func (r *RepoProduct) GetByID(id int) (Product, error) {
	var product Product

	err := r.db.Connect.QueryRowx("SELECT id, name, price FROM products WHERE id=$1", id).StructScan(&product)
	if errors.Is(err, sql.ErrNoRows) {
		return product, fmt.Errorf("not found product {%d}: %w", id, db.ErrDBNotFound)
	}
	if err != nil {
		return product, fmt.Errorf("select product {%d}: %w", id, err)
	}

	return product, nil
}

func (r *RepoProduct) Update(product Product) error {
	result, err := r.db.Connect.NamedExec("UPDATE products SET name=:name, price=:price WHERE id=:id", product)
	if err != nil {
		if r.db.IsErrDuplicate(err) {
			return fmt.Errorf("already exist product with same name {%v}: %w", product, errors.Join(db.ErrDBDuplicateKey, err))
		}
		return fmt.Errorf("sql update product{%v}: %w", product, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected product {%v}: %w", product, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("update product {%v}: %w", product, db.ErrDBNotFound)
	}

	return nil
}

func (r *RepoProduct) GetManyByProductsIDWithTx(tx *sqlx.Tx, productsID []int) ([]Product, error) {
	query, args, err := sqlx.In("SELECT id, name, price FROM products WHERE id IN (?);", productsID)
	if err != nil {
		return []Product{}, fmt.Errorf("prepare sql select in for products {%v}: %w", productsID, err)
	}

	query = tx.Rebind(query)

	products := []Product{}

	rows, err := tx.Queryx(query, args...)
	if err != nil {
		return []Product{}, fmt.Errorf("run sql select in for products {%v}: %w", productsID, err)
	}

	for rows.Next() {
		var product Product
		err = rows.StructScan(&product)
		if err != nil {
			return []Product{}, fmt.Errorf("mapping product struct {%v}: %w", productsID, err)
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *RepoProduct) GetManyByOrderIDWithTx(tx *sqlx.Tx, orderID int) ([]Product, error) {
	products := []Product{}

	query := "SELECT id, name, price FROM products p JOIN orders_products op ON op.product_id = p.id WHERE op.order_id=$1"
	err := tx.Select(&products, query, orderID)
	if err != nil {
		return []Product{}, fmt.Errorf("select products {%d}: %w", orderID, err)
	}

	return products, nil
}
