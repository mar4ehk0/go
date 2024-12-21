package product

type Service struct {
	repo *RepoProduct
}

func NewService(repo *RepoProduct) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(dto *EntryDto) (Product, error) {
	id, err := s.repo.Add(dto)

	return Product{ID: id, Name: dto.Name, Price: dto.Price}, err
}

func (s *Service) GetByID(id int) (Product, error) {
	product, err := s.repo.GetByID(id)

	return product, err
}

func (s *Service) Update(id int, dto *EntryDto) (Product, error) {
	product := Product{ID: id, Name: dto.Name, Price: dto.Price}

	err := s.repo.Update(product)

	return product, err
}
