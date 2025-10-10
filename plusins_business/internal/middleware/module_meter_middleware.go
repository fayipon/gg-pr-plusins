package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ModuleMeterMiddleware struct {
	redis *redis.Redis
}

func NewModuleMeterMiddleware(rds *redis.Redis) *ModuleMeterMiddleware {
	return &ModuleMeterMiddleware{redis: rds}
}

func (m *ModuleMeterMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next(w, r)

		duration := time.Since(start)
		key := "module:usage:" + r.URL.Path

		_, err := m.redis.Incr(key)
		if err != nil {
			log.Printf("[Meter] Redis incr failed: %v", err)
		}

		log.Printf("[Meter] %s %.2fms", r.URL.Path, float64(duration.Milliseconds()))
	}
}
