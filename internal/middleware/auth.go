package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secretKey := os.Getenv("SECRET_KEY")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusForbidden)
			return
		}

		// First part is Bearer, second has to token
		// TODO: Bearer is not required, any string is okay. Check this later.
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			http.Error(w, "Wrong token", http.StatusForbidden)
			return
		}

		tokenPart := parts[1]
		token, err := jwt.Parse(tokenPart, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"]) //alg: HS256 //jwt.SigningMethodHS256
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			http.Error(w, "Token is not valid", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
}
