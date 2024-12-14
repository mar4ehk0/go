package db

import "errors"

var ErrDBDuplicateKey = errors.New("duplicate key")
var ErrDBNotFound = errors.New("not found")
