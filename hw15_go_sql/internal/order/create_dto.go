package order

import (
	"encoding/json"
	"io"
)

type CreateDto struct {
	UserID     int   `json:"user_id"`     //nolint:all
	ProductsID []int `json:"products_id"` //nolint:all
}

func NewCreateDto(r io.Reader) (*CreateDto, error) {
	var dto CreateDto
	err := json.NewDecoder(r).Decode(&dto)
	if err != nil {
		return &dto, err
	}

	if dto.UserID < 1 || len(dto.ProductsID) < 1 {
		return &dto, ErrNotValidRequest
	}

	return &dto, nil
}
