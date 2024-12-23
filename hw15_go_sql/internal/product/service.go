package product

import "fmt"

type Service struct {
	repo *RepoProduct
}

func NewService(repo *RepoProduct) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(dto *EntryDto) (Product, error) {
	id, err := s.repo.Add(dto)
	if err != nil {
		return Product{}, fmt.Errorf("repo add: %w", err)
	}
	return Product{ID: id, Name: dto.Name, Price: dto.Price}, err
}

func (s *Service) GetByID(id int) (Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return Product{}, fmt.Errorf("repo GetByID: %w", err)
	}
	return product, err
}

func (s *Service) Update(id int, dto *EntryDto) (Product, error) {
	product := Product{ID: id, Name: dto.Name, Price: dto.Price}

	err := s.repo.Update(product)
	if err != nil {
		return Product{}, fmt.Errorf("repo Update: %w", err)
	}
	return product, err
}
