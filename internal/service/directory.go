package service

import (
	"github.com/lutfifadlan/directories/internal/model"
	"github.com/lutfifadlan/directories/internal/repository"
)

type DirectoryService struct {
	Repo *repository.DirectoryRepository
}

func NewDirectoryService(repo *repository.DirectoryRepository) *DirectoryService {
	return &DirectoryService{
		Repo: repo,
	}
}

func (s *DirectoryService) AddDirectory(name string) (*model.Directory, error) {
	return s.Repo.Add(name)
}

func (s *DirectoryService) GetDirectoryById(id string) (*model.Directory, error) {
	d, err := s.Repo.FindById(id)

	if err != nil {
		if err == repository.ErrDirectoryNotFound {
			return nil, nil
		}

		return nil, err
	}

	return d, nil
}
