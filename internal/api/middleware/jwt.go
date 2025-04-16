package middlewares

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_KEY")) // use env var in prod

type key string

const UserIDKey key = "userID"

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["user_id"].(float64) // jwt stores numbers as float64
		if !ok {
			http.Error(w, "Invalid token payload", http.StatusUnauthorized)
			return
		}

		// Pass userID into context if needed later
		ctx := context.WithValue(r.Context(), UserIDKey, int(userID))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
