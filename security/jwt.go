package security

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}
