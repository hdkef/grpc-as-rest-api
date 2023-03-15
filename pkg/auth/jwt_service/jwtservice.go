package jwtservice

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService struct {
}

// GenerateToken implements JWTService_
func (*JWTService) GenerateToken(userId string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().UTC().Add(1 * 24 * time.Hour).Unix(),
		"userId": userId,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken implements JWTService_
func (*JWTService) ParseToken(tokenString string, secret string) (string, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
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
	GenerateToken(userId string, secret string) (string, error)
	ParseToken(tokenString string, secret string) (string, error)
}

func NewJWTService() JWTService_ {
	return &JWTService{}
}
