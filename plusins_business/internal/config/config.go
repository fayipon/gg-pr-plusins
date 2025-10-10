package config

import (
	sacfg "github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/plugins/casbin"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth               rest.AuthConf
	CROSConf           sacfg.CROSConf
	CasbinDatabaseConf sacfg.DatabaseConf
	RedisConf          sacfg.RedisConf
	CasbinConf         casbin.CasbinConf
	I18nConf           i18n.Conf
	Log                logx.LogConf

	CoreAPI struct {
		Host      string
		JwtSecret string
	}

	Rpc struct {
		Endpoint string
	}

	DatabaseConf sacfg.DatabaseConf
}
