package svc

import (
    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/config"
    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/model"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
    Config          config.Config
    UserModel       model.UsersModel
    UserLevelModel  model.UserLevelsModel
    UserGroupModel model.UserGroupsModel
    UserTagModel     model.UserTagModel
}

func NewServiceContext(c config.Config) *ServiceContext {
    conn := sqlx.NewMysql(c.DB.DataSource)

    return &ServiceContext{
        Config:         c,
        UserModel:      model.NewUsersModel(conn, c.Cache),
        UserLevelModel: model.NewUserLevelsModel(conn, c.Cache),
        UserGroupModel: model.NewUserGroupsModel(conn, c.Cache),
        UserTagModel:   model.NewUserTagModel(conn, c.Cache),
    }
}
