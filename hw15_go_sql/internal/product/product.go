package product

import (
	"encoding/json"
	"errors"
	"io"
)

var ErrNotValid = errors.New("not valid request")

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
		return &dto, err
	}

	if len(dto.Name) < 1 || dto.Price < 1 {
		return &dto, ErrNotValid
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
