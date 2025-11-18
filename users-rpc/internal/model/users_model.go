package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel 自定义接口（你可以在这里扩展更多方法）
	UsersModel interface {
		usersModel                 // 自动生成的方法（在 users_model_gen.go）
	}

	customUsersModel struct {
		*defaultUsersModel        // 自动生成的 struct
	}
)

// NewUsersModel 创建模型实例
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c),
	}
}
