package middleware

import (
	"context"
	"net/http"
	"strings"

	"parkping/internal/auth"
)

type ctxKey string

const UserIDKey ctxKey = "user_id"

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		userID, err := auth.ParseToken(strings.TrimPrefix(h, "Bearer "))
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		if userID == 0 {
			http.Error(w, "invalid user_id", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, uint(userID))
		next(w, r.WithContext(ctx))
	}
}
