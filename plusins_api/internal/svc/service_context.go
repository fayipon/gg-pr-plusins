package svc

import (
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/config"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/middleware"
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config      config.Config
	Casbin      *casbin.Enforcer
	ModuleMeter rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := c.RedisConf.MustNewUniversalRedis()

	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(
		c.DatabaseConf.Type,
		c.DatabaseConf.GetDSN(),
		c.RedisConf,
	)

	return &ServiceContext{
		Config:      c,
		Casbin:      cbn,
		ModuleMeter: middleware.NewModuleMeterMiddleware(cbn, rds).Handle,
	}
}
