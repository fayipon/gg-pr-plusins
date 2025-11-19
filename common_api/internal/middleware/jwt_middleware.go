package middleware

import (
    "net/http"
    "strings"

    "github.com/golang-jwt/jwt/v4"
)

type JwtMiddleware struct {
    Secret string
}

func NewJwtMiddleware(secret string) *JwtMiddleware {
    return &JwtMiddleware{
        Secret: secret,
    }
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "missing Authorization header", http.StatusUnauthorized)
            return
        }

        // Bearer token
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            http.Error(w, "invalid Authorization format", http.StatusUnauthorized)
            return
        }

        tokenStr := parts[1]

        // Parse JWT
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return []byte(m.Secret), nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "invalid token", http.StatusUnauthorized)
            return
        }

        next(w, r)
    }
}
