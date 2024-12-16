package user

import (
	"encoding/json"
	"io"
	"net/mail"
)

type UpdateDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUpdateDto(r io.Reader) (*UpdateDto, error) {
	var dto UpdateDto
	err := json.NewDecoder(r).Decode(&dto)
	if err != nil {
		return &dto, err
	}

	if len(dto.Name) < 1 || len(dto.Email) < 1 {
		return &dto, ErrNotValidRequest
	}

	_, err = mail.ParseAddress(dto.Email)
	if err != nil {
		return &dto, ErrEmailNotValid
	}

	return &dto, nil
}
