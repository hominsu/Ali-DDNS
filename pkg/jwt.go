package pkg

import (
	"Ali-DDNS/pkg/conf"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtKey = []byte(conf.Option().JwtToken())

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func GenToken(username string) (string, error) {
	expireTime := jwt.NewNumericDate(time.Now().Add(time.Hour))
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "127.0.0.1",
			Subject:   "user token",
			ExpiresAt: expireTime,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(*jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, nil, err
	}

	return token, claims, nil
}
