package repositories

import "backend/internal/models"

var users = []models.User{
	{ID: 1, Email: "john@example.com", Password: "123"},
	{ID: 2, Email: "jane@example.com", Password: "#-123%"},
	{ID: 3, Email: "alice@example.com", Password: "&12*1"},
}

type UserRepository struct{}

func (r *UserRepository) GetUsers() (*[]models.User, error) {
	users_temp := users
	return &users_temp, nil
}

func (r *UserRepository) GetUserByID(id int64) (*models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, nil // Trả về nil nếu không tìm thấy
}
