package pkg

import (
	"Ali-DDNS/pkg/conf"
	"context"
	"github.com/golang-jwt/jwt/v4"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
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

func CheckAuth(ctx context.Context) (string, error) {
	authString, err := grpc_auth.AuthFromMD(ctx, "bearer")
	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "bearer" {
		return "", status.Errorf(codes.InvalidArgument, "token invalid")
	}

	tokenString := kv[1]

	token, claims, err := ParseToken(tokenString)
	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, " %v", err)
	}

	if !token.Valid {
		return "", status.Errorf(codes.Unauthenticated, "Unauthorized")
	}

	return claims.Username, nil
}
