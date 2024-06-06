package middleware

import (
	"context"
	"errors"
	"hello/internal/lib/jwt"
	res "hello/packages/http"
	"net/http"
	"strings"
)

func BasicAuth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				res.Response(w, res.Error(4030, "Authorization header is required"), http.StatusForbidden)
				return
			}

			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
			claims, err := jwt.ParseToken(tokenString, jwt.SigningKey)
			if err != nil {
				if errors.Is(err, jwt.ErrTokenExpired) {
					res.Response(w, res.Error(4010, "Token expired"), http.StatusUnauthorized)
					return
				}
				res.Response(w, res.Error(4010, "Invalid authorization token"), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "uid", claims.UID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
