package entities

import (
	"github.com/dgrijalva/jwt-go"
)

// User entity is user of application
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// SigninCredentials - User sent credential
type SigninCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthObject - Return of Signin API
type AuthObject struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// CustomClaimsJWT - For all important information realted to JWT
type CustomClaimsJWT struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
