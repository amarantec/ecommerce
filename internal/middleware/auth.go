package middleware

import (
  "context"
  "net/http"

  "github.com/amarantec/ecommerce/internal/utils"
)

type contextKey string

const UserIdKey contextKey = "UserId"

func Authenticate (next http.HandlerFunc) http.HandlerFunc {
  return (w http.ResponseWriter, r *http.Request) {
    token := r.Header.Get("Authorization")
    if token == "" {
      http.Error(w, "token is empty", http.StatusUnauthorized)
      return
    }

    userId, err := utils.VerifyToken(token)
    if err != nil {
      http.Error(w, err.Error(), http.StatusUnauthorized)
      return
    }

    ctx := context.WithValue(r.Context(), UserIdKey, userId)
    next(w, r.WithContext(ctx))
  }  
} 
