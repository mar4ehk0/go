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

func (r *Repo) Add(dto CreateDto) (int, error) {
	stmt, err := r.db.PrepareNamed("INSERT INTO products (name, price) VALUES (:name, :price) RETURNING id")
	if err != nil {
		wrappedErr := fmt.Errorf("can't do prepare query product {%s, %d} error: %w", dto.Name, dto.Price, err)
		return 0, wrappedErr
	}

	var id int
	err = stmt.Get(&id, dto)

	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok {
			switch pgErr.Code {
			case "23505":
				return 0, db.ErrDBDuplicateKey
			default:
				wrappedErr := fmt.Errorf("can't do insert product {%s, %d} error: %w", dto.Name, dto.Price, err)
				return 0, wrappedErr
			}
		}
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
