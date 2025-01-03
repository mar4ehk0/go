package order

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/mar4ehk0/go/hw16_docker/internal/product"
	"github.com/mar4ehk0/go/hw16_docker/internal/user"
	"github.com/mar4ehk0/go/hw16_docker/pkg/helper"
)

var (
	ErrUserIDWrongValue = errors.New("userID wrong value")
	ErrProductsIDEmpty  = errors.New("productsID empty")
)

type (
	Order struct {
		ID          int       `json:"id"`
		UserID      int       `db:"user_id" json:"user_id"`           //nolint:all
		OrderDate   time.Time `db:"order_date" json:"order_date"`     //nolint:all
		TotalAmount int       `db:"total_amount" json:"total_amount"` //nolint:all
	}

	OutputReadDto struct {
		ID          int               `json:"id"`
		User        user.User         `json:"user"`
		OrderDate   time.Time         `json:"order_date"`   //nolint:all
		TotalAmount int               `json:"total_amount"` //nolint:all
		Products    []product.Product `json:"product"`      //nolint:all
	}

	EntryCreateDto struct {
		UserID     int   `json:"user_id"`     //nolint:all
		ProductsID []int `json:"products_id"` //nolint:all
	}

	ResponseCreateDto struct {
		ID int `json:"id"`
	}

	EntryUpdateDto struct {
		ProductsID []int `json:"products_id"` //nolint:all
	}

	OutputUpdateDto struct {
		ID          int               `json:"id"`
		OrderDate   time.Time         `json:"order_date"`   //nolint:all
		TotalAmount int               `json:"total_amount"` //nolint:all
		Products    []product.Product `json:"product"`      //nolint:all
	}
)

func NewEntryCreateDto(r io.Reader) (*EntryCreateDto, error) {
	var dto EntryCreateDto
	err := json.NewDecoder(r).Decode(&dto)
	if err != nil {
		return &dto, fmt.Errorf("decode order entry create dto: %w", err)
	}

	if dto.UserID < 1 {
		return &dto, helper.CreateErrorForDto(dto, ErrUserIDWrongValue)
	}

	if len(dto.ProductsID) < 1 {
		return &dto, helper.CreateErrorForDto(dto, ErrProductsIDEmpty)
	}

	return &dto, nil
}

func NewResponseCreateDto(order Order) ([]byte, error) {
	dto := ResponseCreateDto{ID: order.ID}
	data, err := json.Marshal(dto)
	return data, err
}

func NewResponseReadDto(order OutputReadDto) ([]byte, error) {
	data, err := json.Marshal(order)
	return data, err
}

func NewEntryUpdateDto(r io.Reader) (*EntryUpdateDto, error) {
	var dto EntryUpdateDto
	err := json.NewDecoder(r).Decode(&dto)
	if err != nil {
		return &dto, err
	}

	if len(dto.ProductsID) < 1 {
		return &dto, helper.CreateErrorForDto(dto, ErrProductsIDEmpty)
	}

	return &dto, nil
}
