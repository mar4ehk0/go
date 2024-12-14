package product

type CreateDto struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Service struct {
	repo *Repo
}

func NewService(repo *Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(dto CreateDto) (Product, error) {
	id, err := s.repo.Add(dto)

	return Product{Id: id, Name: dto.Name, Price: dto.Price}, err
}

func (s *Service) GetById(id int) (Product, error) {
	product, err := s.repo.GetById(id)

	return product, err
}
