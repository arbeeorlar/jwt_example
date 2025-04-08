package utils

import (
	"github.com/arbeeorlar/jwt-example/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GeneratePasswordhash(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPass), err
}

func ComparePassword(newPassword string, oldPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(newPassword), []byte(oldPassword))
	return err
}

func ParseToken(tokenString string) (claims *models.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)

	if !ok {
		log.Fatal(err)
		return nil, err
	}

	return claims, nil
}
