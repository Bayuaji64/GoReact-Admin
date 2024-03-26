package utility

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID int) (string, error) {

	SecretKey := os.Getenv("JWT_KEY")

	if SecretKey == "" {
		return "", errors.New("JWT_KEY is not set in the environment variables")
	}

	claims := &jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(userID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type Claims struct {
	*jwt.RegisteredClaims
}

func ParseJWT(tokenString string) (*jwt.Token, *Claims, error) {

	SecretKey := os.Getenv("JWT_KEY")

	if SecretKey == "" {
		return nil, nil, errors.New("JWT_KEY is not set in the environment variables")
	}
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, nil, err
	}

	return token, claims, nil
}
