package helper

import (
	"time"

	"github.com/blacheinc/gotemplate/config"
	"github.com/blacheinc/gotemplate/typing"
	"github.com/golang-jwt/jwt/v4"
)

// SignJWT signs a JWT with the given address
func SignJWT(id string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, typing.JWTClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			// expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(2 * time.Hour)),
			Issuer:    config.Env.AppName,
		},
	}).SignedString([]byte(config.Env.JWTSecret))
}
