package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type Config struct {
	Database struct {
		DataSource string
	}
	Cache cache.CacheConf
}
