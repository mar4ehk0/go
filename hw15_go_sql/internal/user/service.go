package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrHashPassword = errors.New("hash password")
	ErrCreate       = errors.New("user create")
	ErrGetByID      = errors.New("user getbyid")
	ErrUpdateByID   = errors.New("user updatebyid")
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
		return User{}, fmt.Errorf("hash password: %w", err)
	}

	dto.Password = string(hashedPassword)
	id, err := s.repo.Add(dto)
	if err != nil {
		return User{}, fmt.Errorf("user repo add: %w", err)
	}

	return User{ID: id, Name: dto.Name, Email: dto.Email, Password: dto.Password}, nil
}

func (s *Service) GetByID(id int) (User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return User{}, fmt.Errorf("user repo getbyid: %w", err)
	}

	return user, err
}

func (s *Service) UpdateByID(id int, dto *EntryUpdateDto) error {
	err := s.repo.Update(id, dto)
	if err != nil {
		return fmt.Errorf("user repo update: %w", err)
	}
	return nil
}
