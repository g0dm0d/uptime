package jwtmanager

import (
	"time"

	"github.com/g0dm0d/uptime/model"
	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessLiveTime = 24 * time.Hour
)

type JWTManager interface {
	GenerateAccessToken(user model.User)
	ValidateJWTToken(token string) (*Claims, error)
}

type Tool struct {
	secret string
}

func New(secret string) *Tool {
	return &Tool{
		secret: secret,
	}
}

type Claims struct {
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
	jwt.RegisteredClaims
}

func (t *Tool) GenerateAccessToken(user model.User) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessLiveTime)),
		},
		Username: user.Username,
		UserID:   user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(t.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *Tool) ValidateJWTToken(token string) (*Claims, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	})
	if err != nil {
		return &Claims{}, err
	}

	if !tkn.Valid {
		return &Claims{}, nil
	}

	return claims, nil
}
