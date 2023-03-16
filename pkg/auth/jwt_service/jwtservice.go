package jwtservice

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService struct {
	SECRET string
}

// GenerateToken implements JWTService_
func (u *JWTService) GenerateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().UTC().Add(1 * 24 * time.Hour).Unix(),
		"userId": userId,
	})

	tokenString, err := token.SignedString([]byte(u.SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken implements JWTService_
func (u *JWTService) ParseToken(tokenString string) (string, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(u.SECRET), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["userId"].(string)
		return userId, nil
	} else {
		return "", errors.New("unauthorized or expired token")
	}
}

type JWTService_ interface {
	GenerateToken(userId string) (string, error)
	ParseToken(tokenString string) (string, error)
}

func NewJWTService(secret string) JWTService_ {
	return &JWTService{
		SECRET: secret,
	}
}
