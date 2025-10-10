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
	rds := redis.MustNewRedis(c.RedisConf) // ✅ 使用 go-zero Redis
	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(
		c.CasbinDatabaseConf.Type,
		c.CasbinDatabaseConf.GetDSN(),
		c.RedisConf,
	)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewModuleMeterMiddleware(cbn, rds).Handle,
	}
}
