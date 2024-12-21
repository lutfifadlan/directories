package model

import "time"

type User struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Role string

const (
	Admin      Role = "admin"
	Maintainer Role = "maintainer"
	Viewer     Role = "viewer"
	Guest      Role = "guest"
)
