package helper

import (
	"github.com/blacheinc/gotemplate/config"
	"github.com/blacheinc/gotemplate/typing"
	"github.com/golang-jwt/jwt/v4"
)

// VerifyJWT verifies a JWT and returns the claims
func VerifyJWT(token string) (*typing.JWTClaims, bool) {
	if twc, err := jwt.ParseWithClaims(token, &typing.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Env.JWTSecret), nil
	}); err == nil && twc.Valid {
		return twc.Claims.(*typing.JWTClaims), true
	}
	return nil, false
}
