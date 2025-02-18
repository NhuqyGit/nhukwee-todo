package services

import (
	"backend/internal/models"
	"backend/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() (*[]models.User, error) {
	return s.repo.GetUsers()
}

func (s *UserService) GetUserByID(id int64) (*models.User, error) {
	return s.repo.GetUserByID(id)
}
