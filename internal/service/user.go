package service

import (
	"github.com/lutfifadlan/directories/internal/model"
	"github.com/lutfifadlan/directories/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) AddUser(email string, role model.Role) (*model.User, error) {
	return s.Repo.Add(email, role)
}

func (s *UserService) GetUserById(id string) (*model.User, error) {
	u, err := s.Repo.FindById(id)

	if err != nil {
		if err == repository.ErrUserNotFound {
			return nil, nil
		}

		return nil, err
	}

	return u, nil
}
