package db

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

var (
	ErrDBDuplicateKey = errors.New("duplicate key")
	ErrDBNotFound     = errors.New("not found")
)

type Connect struct {
	Connect *sqlx.DB
}

func NewDBConnect(db *sqlx.DB) *Connect {
	return &Connect{Connect: db}
}

func (d *Connect) NewTransaction() (*sqlx.Tx, error) {
	tx, err := d.Connect.Beginx()

	return tx, err
}

func (d *Connect) IsErrDuplicate(err error) bool {
	var pgErr pgx.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}

	return false
}

func ProcessError(err error, msgError string) error {
	pgErr, ok := err.(pgx.PgError)

	if ok {
		switch pgErr.Code {
		case "23505":
			return ErrDBDuplicateKey
		default:
			wrappedErr := fmt.Errorf("%s error: %w", msgError, err)
			return wrappedErr
		}
	}

	return err
}
