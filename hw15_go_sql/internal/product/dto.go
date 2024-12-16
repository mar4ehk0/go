package product

import (
	"encoding/json"
	"errors"
	"io"
)

type Dto struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var ErrNotValid = errors.New("not valid request")

func NewDto(r io.Reader) (*Dto, error) {
	var dto Dto
	err := json.NewDecoder(r).Decode(&dto)
	if err != nil {
		return &dto, err
	}

	if len(dto.Name) < 1 || dto.Price < 1 {
		return &dto, ErrNotValid
	}

	return &dto, nil
}
