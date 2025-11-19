package svc

import (
	"github.com/fayipon/gg-pr-plusins/users_rpc/users"
	"github.com/zeromicro/go-zero/zrpc"
	"common_api/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	UsersRpc users.UsersClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UsersRpc: users.NewUsersClient(zrpc.MustNewClient(c.UsersRpc)),
	}
}
