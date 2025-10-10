package middleware

import (
	"net/http"
)

// BusinessAuthMiddleware
// 用于验证商务点固定 Token 与白名单 IP。
type BusinessAuthMiddleware struct{}

func NewBusinessAuthMiddleware() *BusinessAuthMiddleware {
	return &BusinessAuthMiddleware{}
}

func (m *BusinessAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Business-Token")
		ip := r.RemoteAddr

		if token == "" {
			http.Error(w, "Unauthorized Business Token", http.StatusUnauthorized)
			return
		}

		// TODO: 从数据库验证 token 与白名单
		_ = ip

		next(w, r)
	}
}
