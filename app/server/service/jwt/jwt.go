package jwt

import (
	"Ali-DDNS/app/server/service/jwt/conf"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var (
	headerAuthorize = "authorization"
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

// GenToken generate the jwt token
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

// ParseToken parse the jwt token
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

// CheckAuth check the user auth
func CheckAuth(ctx context.Context) (string, error) {
	tokenString := metautils.ExtractIncoming(ctx).Get(headerAuthorize)

	if tokenString == "" {
		return "", status.Errorf(codes.InvalidArgument, "token invalid")
	}

	token, claims, err := ParseToken(tokenString)
	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, " %v", err)
	}

	if !token.Valid {
		return "", status.Errorf(codes.Unauthenticated, "Unauthorized")
	}

	return claims.Username, nil
}
