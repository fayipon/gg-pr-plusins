package svc

import (
    "common_api/internal/config"
	"common_api/internal/middleware"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/zrpc"
    "github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
    Config   config.Config
    UsersRpc users.UsersClient

	JwtMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
    rpcClient := zrpc.MustNewClient(c.UsersRpc)

    return &ServiceContext{
        Config:   c,
        UsersRpc: users.NewUsersClient(rpcClient.Conn()),
		JwtMiddleware: middleware.NewJwtMiddleware(c.JwtAuth.AccessSecret).Handle,
    }
}
