package user

import "errors"

var (
	ErrNotValidRequest = errors.New("not valid request")
	ErrEmailNotValid   = errors.New("not valid email")
)
