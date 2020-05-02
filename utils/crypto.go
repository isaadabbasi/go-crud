package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/isaadabbasi/go_crud/entities"
	"golang.org/x/crypto/bcrypt"
)

// GetEncCost - Get hash roundss
func getEncCost() int {
	return bcrypt.DefaultCost
}

// GetHashedPassword -
func GetHashedPassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), getEncCost())
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// GetSecretKeyJWT - Return Secret Key
func GetSecretKeyJWT() string {
	return "s3cr37-c0d3-bl4h-bl4h"
}

// GenerateToken - Generates HS-256 JWT Token. (KIS)
func GenerateToken() (string, error) {
	claims := entities.CustomClaimsJWT{
		Username: "abcd",
		Email:    "adsasd",
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			Issuer:    "github.com/isaadabbasi",
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(GetSecretKeyJWT()))
	if err != nil {
		return "", errors.New("Error Signing JWT")
	}
	return t, nil
}
