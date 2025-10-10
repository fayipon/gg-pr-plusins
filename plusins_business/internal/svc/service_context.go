package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/config"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config      config.Config
	Casbin      *casbin.Enforcer
	ModuleMeter rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	// ✅ 使用 simple-admin-common 提供的方法初始化 Redis
	rds := c.RedisConf.MustNewUniversalRedis()

	// ✅ 初始化 Casbin 权限系统（带 Redis Watcher）
	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(
		c.DatabaseConf.Type,
		c.DatabaseConf.GetDSN(),
		c.RedisConf,
	)

	return &ServiceContext{
		Config:      c,
		Casbin:      cbn,
		ModuleMeter: middleware.NewModuleMeterMiddleware(rds).Handle,
	}
}
