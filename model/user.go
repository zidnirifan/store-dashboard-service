package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type PayloadAccessToken struct {
	ID    string
	Email string
	Role  string
	jwt.RegisteredClaims
}

type PayloadRefreshToken struct {
	ID    string
	Email string
	jwt.RegisteredClaims
}
