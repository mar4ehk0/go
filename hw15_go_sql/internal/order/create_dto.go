package order

import (
	"encoding/json"
	"io"
)

type CreateDto struct {
	UserId     int   `json:"user_id"`
	ProductsId []int `json:"products_id"`
}

func NewCreateDto(r io.Reader) (*CreateDto, error) {
	var dto CreateDto
	err := json.NewDecoder(r).Decode(&dto)
	if err != nil {
		return &dto, err
	}

	if dto.UserId < 1 || len(dto.ProductsId) < 1 {
		return &dto, ErrNotValidRequest
	}

	return &dto, nil
}
