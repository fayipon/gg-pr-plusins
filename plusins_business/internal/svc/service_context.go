package svc

import (
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/config"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/middleware"
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	Casbin    *casbin.Enforcer
	Authority rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	// ✅ 手动构建 go-zero RedisConf
	redisConf := redis.RedisConf{
		Host: c.RedisConf.Host,
		Pass: c.RedisConf.Pass,
	}

	rds := redis.MustNewRedis(redisConf)

	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(
		c.DatabaseConf.Type,
		c.DatabaseConf.GetDSN(),
		c.RedisConf,
	)

	return &ServiceContext{
		Config:    c,
		Casbin:    cbn,
		Authority: middleware.NewModuleMeterMiddleware(rds).Handle, // ✅ 只传 rds
	}
}
