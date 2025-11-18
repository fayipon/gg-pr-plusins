package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Database struct {
		Driver   string
		Host     string
		Port     int
		DBName   string
		Username string
		Password string
	}
}
