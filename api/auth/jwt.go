package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	ExpDuration = time.Minute * 1 // 5
)

func NewToken(id, secret string) (string, error) {
	timeNow := time.Now()
	claims := jwt.StandardClaims{
		Id:        id,
		IssuedAt:  timeNow.Unix(),
		NotBefore: timeNow.Unix(),
		ExpiresAt: timeNow.Add(ExpDuration).Unix(),
	}
	// TODO: CSRF protection
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	return token.SignedString([]byte(secret))
}

func keyFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(secret))
		return b, nil
	}
}
