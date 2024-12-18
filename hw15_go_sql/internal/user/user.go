package user

import (
	"encoding/json"
	"io"
	"net/mail"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EntryCreateDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewEntryCreateDto(r io.Reader) (*EntryCreateDto, error) {
	var dto EntryCreateDto
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

type ResponseCreateDto struct {
	ID int `json:"id"`
}

func NewResponseCreateDto(user User) ([]byte, error) {
	dto := ResponseCreateDto{ID: user.ID}
	data, err := json.Marshal(dto)
	return data, err
}
