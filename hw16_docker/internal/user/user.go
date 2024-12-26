package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/mail"

	"github.com/mar4ehk0/go/hw16_docker/pkg/helper"
)

var (
	ErrNotValidRequest = errors.New("not valid request")
	ErrEmptyName       = errors.New("empty name")
	ErrEmptyEmail      = errors.New("empty email")
	ErrEmptyPassword   = errors.New("empty password")
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
		return &dto, fmt.Errorf("decode user entry create dto: %w", err)
	}

	if len(dto.Name) < 1 {
		return &dto, helper.CreateErrorForDto(dto, ErrEmptyName)
	}

	if len(dto.Email) < 1 {
		return &dto, helper.CreateErrorForDto(dto, ErrEmptyEmail)
	}

	if len(dto.Password) < 1 {
		return &dto, helper.CreateErrorForDto(dto, ErrEmptyPassword)
	}

	_, err = mail.ParseAddress(dto.Email)
	if err != nil {
		return &dto, helper.CreateErrorForDto(dto, err)
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
		return &dto, fmt.Errorf("decode entry update dto %w", err)
	}

	if len(dto.Name) < 1 {
		return &dto, helper.CreateErrorForDto(dto, ErrEmptyName)
	}
	if len(dto.Email) < 1 {
		return &dto, helper.CreateErrorForDto(dto, ErrEmptyEmail)
	}

	_, err = mail.ParseAddress(dto.Email)
	if err != nil {
		return &dto, helper.CreateErrorForDto(dto, err)
	}

	return &dto, nil
}
