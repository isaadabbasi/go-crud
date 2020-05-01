package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/isaadabbasi/go_crud/utils"
)

// ValidateToken - Middleware fn for valdiating JWT token
func ValidateToken(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized"))
			return
		}

		token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
			// Signing method validation
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(utils.GetSecretKeyJWT()), nil
		})

		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized"))
			return
		}
		if token.Valid {
			handler.ServeHTTP(w, r)
		}
	}
}
