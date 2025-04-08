package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	User JWTUser `json:"user"`
	Role string  `json:"role"`
	jwt.StandardClaims
}
