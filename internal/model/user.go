package model

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}

type Role string

const (
	Admin      Role = "admin"
	Maintainer Role = "maintainer"
	Viewer     Role = "viewer"
	Guest      Role = "guest"
)
