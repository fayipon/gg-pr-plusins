package svc

import (
	"common_api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type ServiceContext struct {
	Config   config.Config
	UsersRpc users.UsersClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UsersRpc: users.NewUsersClient(zrpc.MustNewClient(c.UsersRpc).Conn()),
	}
}
