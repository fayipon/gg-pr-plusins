package config

import (
	sacfg "github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/plugins/casbin"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

// ✅ 自定义 AuthConf，等同 go-zero 的旧版配置结构
type AuthConf struct {
	AccessSecret string
	AccessExpire int64
}

// ✅ Plusins Business Service Config
type Config struct {
	rest.RestConf
	Auth               AuthConf               // ✅ 改成自定义结构
	CROSConf           sacfg.CROSConf
	CasbinDatabaseConf sacfg.DatabaseConf
	RedisConf          sacfg.RedisConf
	CasbinConf         casbin.CasbinConf
	Log                logx.LogConf           // ✅ 保留日志设置

	CoreAPI struct {
		Host      string
		JwtSecret string
	}

	Rpc struct {
		Endpoint string
	}

	DatabaseConf sacfg.DatabaseConf
}
