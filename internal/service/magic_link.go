package service

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"time"

	"github.com/lutfifadlan/directories/internal/model"
	"github.com/lutfifadlan/directories/internal/repository"
)

const TokenLength = 32

var ErrMagicLinkExpired = errors.New("magic link expired")

type MagicLinkService struct {
	MagicLinkRepo *repository.MagicLinkRepository
	EmailService  *EmailService
}

func NewMagicLinkService(magicLinkRepo *repository.MagicLinkRepository) *MagicLinkService {
	return &MagicLinkService{
		MagicLinkRepo: magicLinkRepo,
		EmailService:  NewEmailService(),
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

	magicLink, err := s.MagicLinkRepo.Create(user.Id, token, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, err
	}

	if err := s.SendMagicLinkEmail(email, token); err != nil {
		return nil, err
	}

	return magicLink, nil
}

func (s *MagicLinkService) VerifyMagicLink(db *sql.DB, token string) (*model.MagicLink, error) {
	magicLink, err := s.MagicLinkRepo.FindByToken(token)
	if err != nil {
		if err == repository.ErrMagicLinkNotFound {
			return nil, nil
		}
		return nil, err
	}

	if magicLink.ExpiresAt.Before(time.Now()) {
		return nil, ErrMagicLinkExpired
	}

	if magicLink.Used {
		return nil, errors.New("magic link already used")
	}

	magicLink.Used = true
	if _, err := s.MagicLinkRepo.Update(magicLink); err != nil {
		return nil, err
	}

	return magicLink, nil
}

func (s *MagicLinkService) SendMagicLinkEmail(email string, token string) error {
	err := s.EmailService.SendEmail(
		email,
		"Magic Link Verification",
		"Click this link to login: http://localhost:8080/api/magic-links/"+token,
	)

	if err != nil {
		return err
	}

	return nil
}
