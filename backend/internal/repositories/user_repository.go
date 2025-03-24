package repositories

import (
	"backend/internal/database"
	"backend/internal/models"
	"database/sql"
)

// var users = []models.User{
// 	{ID: 1, Username: "aaa", Email: "john@example.com", Password: "123"},
// 	{ID: 2, Username: "aaa", Email: "jane@example.com", Password: "#-123%"},
// 	{ID: 3, Username: "aaa", Email: "alice@example.com", Password: "&12*1"},
// }

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: database.GetDB()}
}

func (r *UserRepository) GetUsers() (*[]models.User, error) {
	query := "SELECT id, username, password FROM Users"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil
}

func (r *UserRepository) GetUserByID(id int64) (*models.User, error) {
	query := "SELECT id, username, email, password FROM Users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByName(username string) (*models.User, error) {
	query := "SELECT id, username, email, password FROM Users WHERE username = $1"
	row := r.db.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO Users (username, email, password) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password)

	return err
}
