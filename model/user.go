package model

import "time"

type User struct {
	ID        string
	Email     string
	Name      string
	Password  string
	Role      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "dashboard_users"
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken  string
	RefreshToken string
}
