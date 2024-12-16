package product

type Service struct {
	repo *Repo
}

func NewService(repo *Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(dto Dto) (Product, error) {
	id, err := s.repo.Add(dto)

	return Product{Id: id, Name: dto.Name, Price: dto.Price}, err
}

func (s *Service) GetById(id int) (Product, error) {
	product, err := s.repo.GetById(id)

	return product, err
}

func (s *Service) Update(id int, dto Dto) (Product, error) {
	product := Product{Id: id, Name: dto.Name, Price: dto.Price}

	err := s.repo.Update(product)

	return product, err
}
