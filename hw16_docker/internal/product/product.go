package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/mar4ehk0/go/hw16_docker/pkg/helper"
)

var (
	ErrEmptyName  = errors.New("empty name")
	ErrEmptyPrice = errors.New("empty price")
)

type (
	Product struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	EntryDto struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	ResponseCreateDto struct {
		ID int `json:"id"`
	}

	ResponseReadDto struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
	}
)

func NewEntryDto(r io.Reader) (*EntryDto, error) {
	var dto EntryDto
	err := json.NewDecoder(r).Decode(&dto)
	if err != nil {
		return &dto, fmt.Errorf("decode product entry create dto: %w", err)
	}

	if len(dto.Name) < 1 {
		return &dto, helper.CreateErrorForDto(dto, ErrEmptyName)
	}
	if dto.Price < 1 {
		return &dto, helper.CreateErrorForDto(dto, ErrEmptyPrice)
	}

	return &dto, nil
}

func NewResponseCreateDto(product Product) ([]byte, error) {
	dto := ResponseCreateDto{ID: product.ID}
	data, err := json.Marshal(dto)
	return data, err
}

func NewResponseReadDto(product Product) ([]byte, error) {
	dto := ResponseReadDto(product)
	data, err := json.Marshal(dto)
	return data, err
}
