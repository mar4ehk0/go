package db

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx"
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
