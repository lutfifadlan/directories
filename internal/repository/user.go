package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/lutfifadlan/directories/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}

var ErrUserNotFound = errors.New("user not found")

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repo *UserRepository) Add(email string, role model.Role) (*model.User, error) {
	id := uuid.New().String()
	query := "INSERT INTO users (id, email, role) VALUES (?, ?, ?)"

	_, err := repo.DB.Exec(query, id, email, role)
	if err != nil {
		log.Println("Error adding user:", err)
		return nil, err
	}

	u, err := repo.FindById(id)
	if err != nil {
		log.Println("Error finding user by ID:", err)
		return nil, err
	}

	return u, nil
}

func (repo *UserRepository) FindById(id string) (*model.User, error) {
	var user model.User
	query := "SELECT * FROM users WHERE id = ?"

	err := repo.DB.QueryRow(query, id).Scan(&user.Id, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println("Error finding user by ID:", err)
		return nil, ErrUserNotFound
	}

	return &user, nil
}
