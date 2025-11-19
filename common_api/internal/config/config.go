package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type JwtAuth struct {
    AccessSecret string `json:"AccessSecret,optional"`
    AccessExpire int64  `json:"AccessExpire,optional"`
}

type Config struct {
	rest.RestConf
	UsersRpc zrpc.RpcClientConf
    JwtAuth  JwtAuth
}
