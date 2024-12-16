package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repo
}

func NewService(repo *Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(dto *CreateDto) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 14)
	if err != nil {
		wrappedErr := fmt.Errorf("can't hashed password error: %w", err)
		return User{}, wrappedErr
	}

	dto.Password = string(hashedPassword)
	id, err := s.repo.Add(dto)

	return User{Id: id, Name: dto.Name, Email: dto.Email, Password: dto.Password}, err
}

func (s *Service) GetById(id int) (User, error) {
	user, err := s.repo.GetById(id)

	return user, err
}

func (s *Service) UpdateById(id int, dto *UpdateDto) error {
	err := s.repo.Update(id, dto)

	return err
}
