package db

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

var ErrDBDuplicateKey = errors.New("duplicate key")
var ErrDBNotFound = errors.New("not found")

func ProcessError(err error, msgError string) error {
	if pgErr, ok := err.(pgx.PgError); ok {
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

func NewTransaction(db *sqlx.DB) (*sqlx.Tx, error) {
	tx, err := db.Beginx()

	return tx, err
}
