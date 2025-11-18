package svc

import (
	"github.com/fayipon/gg-pr-plusins/users_rpc/internal/config"
	"github.com/fayipon/gg-pr-plusins/users_rpc/internal/model"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UsersModel: model.NewUsersModel(c.Mysql.DataSource),
	}
}
