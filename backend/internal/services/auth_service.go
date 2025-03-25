package services

import (
	"backend/internal/models"
	repo "backend/internal/repositories"
	"backend/internal/utils"
)

type AuthService struct {
	userRepo *repo.UserRepository
}

func NewAuthService(userRepo *repo.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Signup(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.userRepo.CreateUser(user)
}

func (s *AuthService) Signin(username, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByName(username)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPassword(user.Password, password) {
		return nil, err
	}

	return user, nil
}
