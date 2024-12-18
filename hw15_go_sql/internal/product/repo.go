package product

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mar4ehk0/go/hw15_go_sql/pkg/db"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Add(dto *EntryDto) (int, error) {
	query := "INSERT INTO products (name, price) VALUES (:name, :price) RETURNING id"
	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		wrappedErr := fmt.Errorf("can't do prepare query product {%s, %d} error: %w", dto.Name, dto.Price, err)
		return 0, wrappedErr
	}

	var id int
	err = stmt.Get(&id, dto)
	if err != nil {
		msgErr := fmt.Sprintf("can't do insert product {%s, %d}", dto.Name, dto.Price)
		err = db.ProcessError(err, msgErr)
		return 0, err
	}

	return id, nil
}

func (r *Repo) GetByID(id int) (Product, error) {
	var product Product

	err := r.db.QueryRowx("SELECT id, name, price FROM products WHERE id=$1", id).StructScan(&product)
	if errors.Is(err, sql.ErrNoRows) {
		return product, db.ErrDBNotFound
	}
	if err != nil {
		wrappedErr := fmt.Errorf("can't do select product by id {%d} error: %w", id, err)
		return product, wrappedErr
	}

	return product, nil
}

func (r *Repo) Update(product Product) error {
	msgErr := fmt.Sprintf("can't do prepare update product {%s, %d}", product.Name, product.Price)

	result, err := r.db.NamedExec("UPDATE products SET name=:name, price=:price WHERE id=:id", product)
	if err != nil {
		err = db.ProcessError(err, msgErr)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err = db.ProcessError(err, msgErr)
		return err
	}
	if rowsAffected == 0 {
		return db.ErrDBNotFound
	}

	return nil
}

func (r *Repo) GetManyWithTx(tx *sqlx.Tx, productsID []int) ([]Product, error) {
	query, args, err := sqlx.In("SELECT id, name, price FROM products WHERE id IN (?);", productsID)
	if err != nil {
		return []Product{}, err
	}

	query = tx.Rebind(query)

	products := []Product{}

	rows, err := tx.Queryx(query, args...)
	if err != nil {
		return []Product{}, err
	}

	for rows.Next() {
		var product Product
		err = rows.StructScan(&product)
		if err != nil {
			return []Product{}, err
		}
		products = append(products, product)
	}

	return products, nil
}
