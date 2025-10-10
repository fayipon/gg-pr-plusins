package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
)

type ModuleMeterMiddleware struct {
	cbn *casbin.Enforcer
	rds redis.UniversalClient
}

func NewModuleMeterMiddleware(cbn *casbin.Enforcer, rds redis.UniversalClient) *ModuleMeterMiddleware {
	return &ModuleMeterMiddleware{
		cbn: cbn,
		rds: rds,
	}
}

func (m *ModuleMeterMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ✅ 这里可以加模块计费或 API 统计逻辑
		next(w, r)
	}
}
