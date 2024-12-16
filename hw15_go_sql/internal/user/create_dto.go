package user

import (
	"encoding/json"
	"io"
	"net/mail"
)

type CreateDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewCreateDto(r io.Reader) (*CreateDto, error) {
	var dto CreateDto
	err := json.NewDecoder(r).Decode(&dto)
	if err != nil {
		return &dto, err
	}

	if len(dto.Name) < 1 || len(dto.Email) < 1 || len(dto.Password) < 6 {
		return &dto, ErrNotValidRequest
	}

	_, err = mail.ParseAddress(dto.Email)
	if err != nil {
		return &dto, ErrEmailNotValid
	}

	return &dto, nil
}
