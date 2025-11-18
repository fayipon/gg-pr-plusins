package svc

import (
	"users-rpc/internal/config"
	"users-rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.Database.DataSource)

	return &ServiceContext{
		Config:     c,
		UsersModel: model.NewUsersModel(conn, c.Cache),
	}
}
