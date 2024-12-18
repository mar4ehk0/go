package order

import (
	"fmt"
	"log"
	"time"

	"github.com/mar4ehk0/go/hw15_go_sql/internal/product"
	"github.com/mar4ehk0/go/hw15_go_sql/internal/user"
	"github.com/mar4ehk0/go/hw15_go_sql/pkg/db"
)

type Service struct {
	repoOrder   *Repo
	repoProduct *product.Repo
	repoUser    *user.Repo
}

func NewService(repoOrder *Repo, repoProduct *product.Repo, repoUser *user.Repo) *Service {
	return &Service{
		repoOrder:   repoOrder,
		repoProduct: repoProduct,
		repoUser:    repoUser,
	}
}

func (s *Service) Create(dto *CreateDto) (Order, error) {
	tx, err := db.NewTransaction(s.repoOrder.db)
	if err != nil {
		return Order{}, err
	}
	defer func() {
		var dbErr error
		if err != nil {
			dbErr = tx.Rollback()
		} else {
			dbErr = tx.Commit()
		}
		if dbErr != nil {
			log.Println(dbErr)
		}
	}()

	user, err := s.repoUser.GetByINWithTx(tx, dto.UserId)
	if err != nil {
		return Order{}, err
	}

	products, err := s.repoProduct.GetByINWithTx(tx, dto.ProductsId)
	if err != nil {
		return Order{}, err
	}

	if len(products) != len(dto.ProductsId) {
		return Order{}, fmt.Errorf("some products with ids %v not found", dto.ProductsId)
	}

	totalAmount := 0
	for _, v := range products {
		totalAmount += v.Price
	}

	orderDate := time.Now()

	id, err := s.repoOrder.Add(tx, user, products, totalAmount, orderDate)

	if err != nil {
		return Order{}, nil
	}

	return Order{Id: id, UserId: user.Id, OrderDate: orderDate, TotalAmount: totalAmount}, err
}
