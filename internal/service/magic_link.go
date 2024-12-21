package service

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"time"

	"github.com/lutfifadlan/directories/internal/model"
	"github.com/lutfifadlan/directories/internal/repository"
)

const TokenLength = 32

type MagicLinkService struct {
	Repo *repository.MagicLinkRepository
}

func NewMagicLinkService(repo *repository.MagicLinkRepository) *MagicLinkService {
	return &MagicLinkService{
		Repo: repo,
	}
}

func GenerateToken() (string, error) {
	bytes := make([]byte, TokenLength)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func (s *MagicLinkService) GenerateMagicLink(db *sql.DB, email string) (*model.MagicLink, error) {
	token, err := GenerateToken()
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository(db)
	user, err := userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	magicLink, err := s.Repo.Create(user.Id, token, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, err
	}

	return magicLink, nil
}
