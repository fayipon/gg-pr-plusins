package model

import (
    "context"
    "fmt"

    "github.com/zeromicro/go-zero/core/stores/cache"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
    UsersModel interface {
        usersModel
        List(ctx context.Context, offset, limit int) ([]*Users, error)
        Count(ctx context.Context) (int64, error)
    }

    customUsersModel struct {
        *defaultUsersModel
    }
)

func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
    return &customUsersModel{
        defaultUsersModel: newUsersModel(conn, c, opts...),
    }
}

// 分页查询
func (m *customUsersModel) List(ctx context.Context, offset, limit int) ([]*Users, error) {
    query := fmt.Sprintf("SELECT %s FROM %s LIMIT ?, ?", usersRows, m.table)

    var list []*Users
    err := m.QueryRowsNoCacheCtx(ctx, &list, query, offset, limit)
    if err != nil {
        return nil, err
    }
    return list, nil
}

// 统计总数
func (m *customUsersModel) Count(ctx context.Context) (int64, error) {
    query := fmt.Sprintf("SELECT COUNT(*) FROM %s", m.table)

    var total int64
    err := m.QueryRowNoCacheCtx(ctx, &total, query)
    if err != nil {
        return 0, err
    }
    return total, nil
}
