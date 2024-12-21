package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *RepoUser
}

func NewService(repo *RepoUser) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(dto *EntryCreateDto) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 14)
	if err != nil {
		wrappedErr := fmt.Errorf("can't hashed password error: %w", err)
		return User{}, wrappedErr
	}

	dto.Password = string(hashedPassword)
	id, err := s.repo.Add(dto)

	return User{ID: id, Name: dto.Name, Email: dto.Email, Password: dto.Password}, err
}

func (s *Service) GetByID(id int) (User, error) {
	user, err := s.repo.GetByID(id)

	return user, err
}

func (s *Service) UpdateByID(id int, dto *EntryUpdateDto) error {
	err := s.repo.Update(id, dto)

	return err
}
