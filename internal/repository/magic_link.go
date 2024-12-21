package repository

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lutfifadlan/directories/internal/model"
)

type MagicLinkRepository struct {
	DB *sql.DB
}

var ErrMagicLinkNotFound = errors.New("magic link not found")

func NewMagicLinkRepository(db *sql.DB) *MagicLinkRepository {
	return &MagicLinkRepository{
		DB: db,
	}
}

func (repo *MagicLinkRepository) Create(userId string, token string, expiresAt time.Time) (*model.MagicLink, error) {
	id := uuid.New().String()
	query := "INSERT INTO magic_links (id, user_id, token, expires_at) VALUES (?, ?, ?, ?)"

	_, err := repo.DB.Exec(query, id, userId, token, expiresAt)
	if err != nil {
		log.Println("Error creating magic link:", err)
		return nil, err
	}

	m, err := repo.FindById(id)
	if err != nil {
		log.Println("Error finding magic link by ID:", err)
		return nil, err
	}

	return m, nil
}

func (repo *MagicLinkRepository) FindById(id string) (*model.MagicLink, error) {
	query := "SELECT * FROM magic_links WHERE id = ?"

	var m model.MagicLink

	err := repo.DB.QueryRow(query, id).Scan(
		&m.ID,
		&m.UserID,
		&m.Token,
		&m.Used,
		&m.ExpiresAt,
		&m.CreatedAt,
		&m.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrMagicLinkNotFound
		}

		log.Println("Error finding magic link by ID:", err)
		return nil, err
	}

	return &m, nil
}
