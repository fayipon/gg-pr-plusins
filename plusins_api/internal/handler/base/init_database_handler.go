package base

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"

	"github.com/fayipon/gg-pr-plusins/plusins_api/internal/config"
	"github.com/fayipon/gg-pr-plusins/plusins_api/internal/middleware"
)

type ServiceContext struct {
	Config    config.Config
	Casbin    *casbin.Enforcer
	Authority rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := c.RedisConf.MustNewUniversalRedis()
	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(
		c.CasbinDatabaseConf.Type,
		c.CasbinDatabaseConf.GetDSN(),
		c.RedisConf,
	)

	return &ServiceContext{
		Config:    c,
		Casbin:    cbn,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds).Handle,
	}
}
