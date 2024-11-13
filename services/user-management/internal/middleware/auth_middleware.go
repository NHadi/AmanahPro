package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

func JWTAuthMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			logrus.Infof("Authorization header received: %s", authHeader) // Log the header for debugging

			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				logrus.Warn("Unauthorized access: Missing or malformed Authorization header")
				http.Error(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
				return
			}

			// Extract token and validate with jwtSecret
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					logrus.Errorf("Unexpected signing method: %v", token.Header["alg"])
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(jwtSecret), nil
			})

			if err != nil || !token.Valid {
				logrus.WithError(err).Warn("Invalid JWT token")
				http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
				return
			}

			logrus.Info("Token validated successfully")
			next.ServeHTTP(w, r)
		})
	}
}
