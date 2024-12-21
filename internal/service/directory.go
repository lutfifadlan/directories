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
