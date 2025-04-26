package service

import (
	"errors"

	"github.com/danielwetan/gin-clean-architecture/internal/model"
	"github.com/danielwetan/gin-clean-architecture/internal/repository"

	"github.com/google/uuid"
)

type UserService interface {
	RegisterUser(name, email string) (*model.User, error)
	GetUser(id string) (*model.User, error)
	ListUsers() ([]*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (s *userService) RegisterUser(name, email string) (*model.User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	user := &model.User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}

	err := s.userRepo.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUser(id string) (*model.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) ListUsers() ([]*model.User, error) {
	return s.userRepo.FindAll()
}
