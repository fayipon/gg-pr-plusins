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
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte(`{"code":401, "message":"missing Authorization"}`))
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte(`{"code":401, "message":"invalid Authorization format"}`))
            return
        }

        tokenStr := parts[1]

        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return []byte(m.Secret), nil
        })

        if err != nil || !token.Valid {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte(`{"code":401, "message":"invalid token"}`))
            return
        }

        // JWT OK → go to handler
        next(w, r)
    }
}
