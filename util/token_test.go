package util

import (
	"store-dashboard-service/config"
	"store-dashboard-service/model"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAccessToken(t *testing.T) {
	user := model.PayloadAccessToken{
		ID:    "1",
		Email: "test@mail.com",
		Role:  "admin",
	}
	token, err := GenerateAccessToken(&user)

	claim := &model.PayloadAccessToken{}
	res, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().AccessTokenKey), nil
	})

	expirationTime := 24 * time.Hour

	assert.Nil(t, err)
	assert.Equal(t, true, res.Valid)
	assert.Equal(t, user.ID, claim.ID)
	assert.Equal(t, user.Email, claim.Email)
	assert.Equal(t, user.Role, claim.Role)
	assert.Equal(t, claim.IssuedAt.Add(expirationTime), claim.ExpiresAt.Time)
}

func TestGenerateRefreshToken(t *testing.T) {
	user := model.PayloadRefreshToken{
		ID:    "1",
		Email: "test@mail.com",
	}
	token, err := GenerateRefreshToken(&user)

	claim := &model.PayloadRefreshToken{}
	res, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().RefreshTokenKey), nil
	})

	expirationTime := 30 * 24 * time.Hour

	assert.Nil(t, err)
	assert.Equal(t, true, res.Valid)
	assert.Equal(t, user.ID, claim.ID)
	assert.Equal(t, user.Email, claim.Email)
	assert.Equal(t, claim.IssuedAt.Add(expirationTime), claim.ExpiresAt.Time)
}
