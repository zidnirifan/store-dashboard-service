package util

import (
	"store-dashboard-service/config"
	"store-dashboard-service/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var accessTokenKey = []byte(config.GetConfig().AccessTokenKey)
var refreshTokenKey = []byte(config.GetConfig().RefreshTokenKey)

func GenerateAccessToken(user *model.PayloadAccessToken) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
			"iat":   time.Now().Unix(),
			"exp":   expirationTime.Unix(),
		},
	)

	return token.SignedString(accessTokenKey)
}

func GenerateRefreshToken(user *model.PayloadRefreshToken) (string, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    user.ID,
			"email": user.Email,
			"iat":   time.Now().Unix(),
			"exp":   expirationTime.Unix(),
		},
	)

	return token.SignedString(refreshTokenKey)
}
