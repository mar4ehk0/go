package user

import (
	"encoding/json"
	"errors"
	"io"
	"net/mail"
)

var (
	ErrNotValidRequest = errors.New("not valid request")
	ErrEmailNotValid   = errors.New("not valid email")
)

type (
	User struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	EntryCreateDto struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	ResponseCreateDto struct {
		ID int `json:"id"`
	}

	ResponseReadDto struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	EntryUpdateDto struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

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

func NewResponseCreateDto(user User) ([]byte, error) {
	dto := ResponseCreateDto{ID: user.ID}
	data, err := json.Marshal(dto)
	return data, err
}

func NewResponseReadDto(user User) ([]byte, error) {
	dto := &ResponseReadDto{ID: user.ID, Name: user.Name, Email: user.Email}
	data, err := json.Marshal(dto)
	return data, err
}

func NewEntryUpdateDto(r io.Reader) (*EntryUpdateDto, error) {
	var dto EntryUpdateDto
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
