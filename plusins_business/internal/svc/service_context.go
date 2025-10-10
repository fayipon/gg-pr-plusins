package svc

import (
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/config"
	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
)

// 全局上下文
type ServiceContext struct {
	Config config.Config
	Casbin *casbin.Enforcer

	BusinessAuth rest.Middleware
	ModuleMeter  rest.Middleware
}

// 初始化依赖注入
func NewServiceContext(c config.Config) *ServiceContext {
	rds := c.RedisConf.MustNewUniversalRedis()

	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(
		c.CasbinDatabaseConf.Type,
		c.CasbinDatabaseConf.GetDSN(),
		c.RedisConf,
	)

	return &ServiceContext{
		Config:        c,
		Casbin:        cbn,
		BusinessAuth:  middleware.NewBusinessAuthMiddleware().Handle,
		ModuleMeter:   middleware.NewModuleMeterMiddleware(rds).Handle,
	}
}
