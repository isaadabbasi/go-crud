package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   "verficiation",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(GetSecretKeyJWT()))
	if err != nil {
		return "", errors.New("Error Signing JWT")
	}
	return t, nil
}
