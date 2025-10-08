package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/fayipon/gg-pr-plusins/plusins_api/internal/config"
	"github.com/fayipon/gg-pr-plusins/plusins_api/internal/middleware"
)

type ServiceContext struct {
	Config    config.Config
	Casbin    *casbin.Enforcer
	Authority rest.Middleware
	Trans     *i18n.Translator
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := c.RedisConf.MustNewUniversalRedis()
	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(
		c.CasbinDatabaseConf.Type,
		c.CasbinDatabaseConf.GetDSN(),
		c.RedisConf,
	)

	// ✅ 新增這行：初始化 Translator，避免 nil pointer
	var trans *i18n.Translator
	if c.I18nConf.Dir != "" {
		trans = i18n.NewTranslator(c.I18nConf)
	}

	return &ServiceContext{
		Config:    c,
		Casbin:    cbn,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds).Handle,
		Trans:     trans,
	}
}
