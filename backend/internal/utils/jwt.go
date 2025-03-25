package utils

import (
	"fmt"
	"time"

	"backend/configs"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(configs.Envs.JWTSecret)

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("method signed token invalid")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
