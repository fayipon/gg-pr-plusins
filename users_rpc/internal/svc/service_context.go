package svc

import (
	"users_rpc/internal/config"
	"users_rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dataSource := fmt.Sprintf("%s://%s:%s@tcp(%s:%d)/%s",
		c.Database.Driver,
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
	)

	conn := sqlx.NewMysql(dataSource)

	return &ServiceContext{
		Config:     c,
		UsersModel: model.NewUsersModel(conn),
	}
}
