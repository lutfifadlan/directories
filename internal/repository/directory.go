package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/lutfifadlan/directories/internal/model"
)

type DirectoryRepository struct {
	DB *sql.DB
}

var ErrDirectoryNotFound = errors.New("directory not found")

func NewDirectoryRepository(db *sql.DB) *DirectoryRepository {
	return &DirectoryRepository{
		DB: db,
	}
}

func (repo *DirectoryRepository) Add(name string) (*model.Directory, error) {
	id := uuid.New().String()
	query := "INSERT INTO directories (id, name) VALUES (?, ?)"

	_, err := repo.DB.Exec(query, id, name)
	if err != nil {
		log.Println("Error adding directory:", err)
		return nil, err
	}

	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return nil, err
	}

	d, err := repo.FindById(id)
	if err != nil {
		log.Println("Error finding directory by ID:", err)
		return nil, err
	}

	return d, nil
}

func (repo *DirectoryRepository) FindAll() ([]model.Directory, error) {
	rows, err := repo.DB.Query("SELECT * FROM directories")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var directories []model.Directory
	for rows.Next() {
		var directory model.Directory
		if err := rows.Scan(&directory.Id, &directory.Name); err != nil {
			log.Println(err)
			return nil, err
		}
		directories = append(directories, directory)
	}

	return directories, nil
}

func (repo *DirectoryRepository) FindById(id string) (*model.Directory, error) {
	query := "SELECT * FROM directories WHERE id = ?"

	var d model.Directory

	err := repo.DB.QueryRow(query, id).Scan(&d.Id, &d.Name, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrDirectoryNotFound
		}
		log.Println("Error finding directory by id:", err)
		return nil, err
	}

	return &d, nil
}
