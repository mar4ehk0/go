package order

import (
	"fmt"
	"time"

	"github.com/mar4ehk0/go/hw15_go_sql/internal/product"
	"github.com/mar4ehk0/go/hw15_go_sql/internal/user"
)

type Service struct {
	repoOrder   *RepoOrder
	repoProduct *product.RepoProduct
	repoUser    *user.RepoUser
}

func NewService(repoOrder *RepoOrder, repoProduct *product.RepoProduct, repoUser *user.RepoUser) *Service {
	return &Service{
		repoOrder:   repoOrder,
		repoProduct: repoProduct,
		repoUser:    repoUser,
	}
}

func (s *Service) Create(dto *EntryCreateDto) (Order, error) {
	tx, err := s.repoOrder.db.NewTransaction()
	if err != nil {
		return Order{}, fmt.Errorf("new transaction: %w", err)
	}
	defer func() {
		var dbErr error
		if err != nil {
			dbErr = tx.Rollback()
			if dbErr != nil {
				fmt.Println(fmt.Errorf("transaction rollback: %w", err))
			}
		} else {
			dbErr = tx.Commit()
			if dbErr != nil {
				fmt.Println(fmt.Errorf("transaction commit: %w", err))
			}
		}
	}()

	user, err := s.repoUser.GetByIDWithTx(tx, dto.UserID)
	if err != nil {
		return Order{}, fmt.Errorf("user repo GetByIDWithTx: %w", err)
	}

	products, err := s.repoProduct.GetManyByProductsIDWithTx(tx, dto.ProductsID)
	if err != nil {
		return Order{}, fmt.Errorf("product repo GetManyByProductsIDWithTx: %w", err)
	}

	if len(products) != len(dto.ProductsID) {
		return Order{}, fmt.Errorf("some products with ids %v not found", dto.ProductsID)
	}

	totalAmount := s.calculateTotalAmount(products)
	orderDate := time.Now()

	id, err := s.repoOrder.AddWithTx(tx, user, products, totalAmount, orderDate)
	if err != nil {
		return Order{}, fmt.Errorf("order repo AddWithTx: %w", err)
	}

	return Order{ID: id, UserID: user.ID, OrderDate: orderDate, TotalAmount: totalAmount}, nil
}

func (s *Service) GetByID(id int) (OutputReadDto, error) {
	tx, err := s.repoOrder.db.NewTransaction()
	if err != nil {
		return OutputReadDto{}, fmt.Errorf("new transaction: %w", err)
	}
	defer func() {
		var dbErr error
		if err != nil {
			dbErr = tx.Rollback()
			if dbErr != nil {
				fmt.Println(fmt.Errorf("transaction rollback: %w", err))
			}
		} else {
			dbErr = tx.Commit()
			if dbErr != nil {
				fmt.Println(fmt.Errorf("transaction commit: %w", err))
			}
		}
	}()

	order, err := s.repoOrder.GetByIDWithTx(tx, id)
	if err != nil {
		return OutputReadDto{}, fmt.Errorf("order repo GetByIDWithTx: %w", err)
	}

	user, err := s.repoUser.GetByIDWithTx(tx, id)
	if err != nil {
		return OutputReadDto{}, fmt.Errorf("user repo GetByIDWithTx: %w", err)
	}

	products, err := s.repoProduct.GetManyByOrderIDWithTx(tx, order.ID)
	if err != nil {
		return OutputReadDto{}, fmt.Errorf("product repo GetManyByOrderIDWithTx: %w", err)
	}

	return OutputReadDto{
		ID:          order.ID,
		User:        user,
		OrderDate:   order.OrderDate,
		TotalAmount: order.TotalAmount,
		Products:    products,
	}, nil
}

func (s *Service) Update(orderID int, dto *EntryUpdateDto) (OutputUpdateDto, error) {
	tx, err := s.repoOrder.db.NewTransaction()
	if err != nil {
		return OutputUpdateDto{}, fmt.Errorf("new transaction: %w", err)
	}
	defer func() {
		var dbErr error
		if err != nil {
			dbErr = tx.Rollback()
			if dbErr != nil {
				fmt.Println(fmt.Errorf("transaction rollback: %w", err))
			}
		} else {
			dbErr = tx.Commit()
			if dbErr != nil {
				fmt.Println(fmt.Errorf("transaction commit: %w", err))
			}
		}
	}()

	order, err := s.repoOrder.GetByIDWithTx(tx, orderID)
	if err != nil {
		return OutputUpdateDto{}, fmt.Errorf("order repo GetByIDWithTx: %w", err)
	}

	products, err := s.repoProduct.GetManyByProductsIDWithTx(tx, dto.ProductsID)
	if err != nil {
		return OutputUpdateDto{}, fmt.Errorf("product repo GetManyByProductsIDWithTx: %w", err)
	}
	if len(products) != len(dto.ProductsID) {
		return OutputUpdateDto{}, fmt.Errorf("some products with ids %v not found", dto.ProductsID)
	}

	err = s.repoOrder.DeleteProductsByIDWithTx(tx, orderID)
	if err != nil {
		return OutputUpdateDto{}, fmt.Errorf("order repo DeleteProductsByIDWithTx: %w", err)
	}

	err = s.repoOrder.AddProductsByIDWithTx(tx, orderID, dto.ProductsID)
	if err != nil {
		return OutputUpdateDto{}, fmt.Errorf("order repo AddProductsByIDWithTx: %w", err)
	}

	totalAmount := s.calculateTotalAmount(products)
	err = s.repoOrder.UpdateTotalAmountOrderWithTX(tx, orderID, totalAmount)
	if err != nil {
		return OutputUpdateDto{}, fmt.Errorf("order repo UpdateTotalAmountOrderWithTX: %w", err)
	}

	return OutputUpdateDto{ID: orderID, OrderDate: order.OrderDate, TotalAmount: totalAmount, Products: products}, err
}

func (s *Service) calculateTotalAmount(products []product.Product) int {
	totalAmount := 0
	for _, v := range products {
		totalAmount += v.Price
	}

	return totalAmount
}
