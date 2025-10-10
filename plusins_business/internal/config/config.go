package config

import (
	sacfg "github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/plugins/casbin"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

// ✅ Plusins Business Service Config
type Config struct {
	rest.RestConf
	Auth               rest.AuthConf
	CROSConf           sacfg.CROSConf
	CasbinDatabaseConf sacfg.DatabaseConf
	RedisConf          sacfg.RedisConf
	CasbinConf         casbin.CasbinConf
	I18nConf           sacfg.I18nConf // ⚠️ 若 i18n 报错，可改为删除此行
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
