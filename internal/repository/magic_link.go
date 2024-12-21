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

	// the order of scan is actually based on the order of the columns in the query
	err := repo.DB.QueryRow(query, id).Scan(
		&m.ID,
		&m.UserID,
		&m.Token,
		&m.Used,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.ExpiresAt,
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

func (repo *MagicLinkRepository) FindByToken(token string) (*model.MagicLink, error) {
	query := "SELECT * FROM magic_links WHERE token = ?"

	var m model.MagicLink

	err := repo.DB.QueryRow(query, token).Scan(
		&m.ID,
		&m.UserID,
		&m.Token,
		&m.Used,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.ExpiresAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrMagicLinkNotFound
		}

		log.Println("Error finding magic link by token:", err)
		return nil, err
	}

	return &m, nil
}

func (repo *MagicLinkRepository) Update(m *model.MagicLink) (*model.MagicLink, error) {
	query := "UPDATE magic_links SET used = ? WHERE id = ?"

	_, err := repo.DB.Exec(query, m.Used, m.ID)
	if err != nil {
		log.Println("Error updating magic link:", err)
		return nil, err
	}

	return m, nil
}
