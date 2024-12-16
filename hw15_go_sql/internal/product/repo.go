package product

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	"github.com/mar4ehk0/go/hw15_go_sql/pkg/db"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Add(dto Dto) (int, error) {
	stmt, err := r.db.PrepareNamed("INSERT INTO products (name, price) VALUES (:name, :price) RETURNING id")
	if err != nil {
		wrappedErr := fmt.Errorf("can't do prepare query product {%s, %d} error: %w", dto.Name, dto.Price, err)
		return 0, wrappedErr
	}

	var id int
	err = stmt.Get(&id, dto)

	if err != nil {
		msgErr := fmt.Sprintf("can't do insert product {%s, %d}", dto.Name, dto.Price)
		err = processError(err, msgErr)
		return 0, err
	}

	return id, nil
}

func (r *Repo) GetById(id int) (Product, error) {
	var product Product

	err := r.db.QueryRowx("SELECT id, name, price FROM products WHERE id=$1", id).StructScan(&product)
	if err == sql.ErrNoRows {
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
		err = processError(err, msgErr)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err = processError(err, msgErr)
		return err
	}
	if rowsAffected == 0 {
		return db.ErrDBNotFound
	}

	return nil
}

func processError(err error, msgError string) error {
	if pgErr, ok := err.(pgx.PgError); ok {
		switch pgErr.Code {
		case "23505":
			return db.ErrDBDuplicateKey
		default:
			wrappedErr := fmt.Errorf("%s error: %w", msgError, err)
			return wrappedErr
		}
	}

	return err
}
