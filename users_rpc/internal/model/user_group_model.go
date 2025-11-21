package model

import (
    "context"
    "database/sql"
    "fmt"

    "github.com/zeromicro/go-zero/core/stores/cache"
    "github.com/zeromicro/go-zero/core/stores/sqlc"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserGroupsModel = (*defaultUserGroupsModel)(nil)

const (
    cacheUserGroupsIdPrefix = "cache:userGroups:id:"

    userGroupsRows = "`id`,`name`,`display_name`,`setting`,`created_at`,`updated_at`"

    userGroupsRowsExpectAutoSet = "`name`,`display_name`,`setting`,`created_at`,`updated_at`"

    userGroupsRowsWithPlaceHolder = "`name`=?,`display_name`=?,`setting`=?,`created_at`=?,`updated_at`=?"
)

type (
    UserGroups struct {
        Id          uint64 `db:"id"`
        Name        string `db:"name"`
        DisplayName string `db:"display_name"`
        Setting     string `db:"setting"`
        CreatedAt   int64  `db:"created_at"`
        UpdatedAt   int64  `db:"updated_at"`
    }

    UserGroupsModel interface {
        Insert(ctx context.Context, data *UserGroups) (sql.Result, error)
        FindOne(ctx context.Context, id uint64) (*UserGroups, error)
        Update(ctx context.Context, data *UserGroups) error
        Delete(ctx context.Context, id uint64) error
        List(ctx context.Context, offset, limit int32) ([]*UserGroups, error)
        Count(ctx context.Context) (int64, error)
    }

    defaultUserGroupsModel struct {
        sqlc.CachedConn
        table string
    }
)

func NewUserGroupsModel(conn sqlx.SqlConn, c cache.CacheConf) UserGroupsModel {
    return &defaultUserGroupsModel{
        CachedConn: sqlc.NewConn(conn, c),
        table:      "`user_groups`",
    }
}

func (m *defaultUserGroupsModel) Insert(ctx context.Context, data *UserGroups) (sql.Result, error) {
    userGroupsIdKey := fmt.Sprintf("%s%v", cacheUserGroupsIdPrefix, data.Id)
    ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userGroupsRowsExpectAutoSet)
        return conn.ExecCtx(ctx, query, data.Name, data.DisplayName, data.Setting, data.CreatedAt, data.UpdatedAt)
    }, userGroupsIdKey)
    return ret, err
}

func (m *defaultUserGroupsModel) FindOne(ctx context.Context, id uint64) (*UserGroups, error) {
    userGroupsIdKey := fmt.Sprintf("%s%v", cacheUserGroupsIdPrefix, id)
    var resp UserGroups
    err := m.QueryRowCtx(ctx, &resp, userGroupsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
        query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userGroupsRows, m.table)
        return conn.QueryRowCtx(ctx, v, query, id)
    })
    switch err {
    case nil:
        return &resp, nil
    case sqlc.ErrNotFound:
        return nil, ErrNotFound
    default:
        return nil, err
    }
}

func (m *defaultUserGroupsModel) Update(ctx context.Context, data *UserGroups) error {
    userGroupsIdKey := fmt.Sprintf("%s%v", cacheUserGroupsIdPrefix, data.Id)
    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userGroupsRowsWithPlaceHolder)
        return conn.ExecCtx(ctx, query, data.Name, data.DisplayName, data.Setting, data.CreatedAt, data.UpdatedAt, data.Id)
    }, userGroupsIdKey)
    return err
}

func (m *defaultUserGroupsModel) Delete(ctx context.Context, id uint64) error {
    userGroupsIdKey := fmt.Sprintf("%s%v", cacheUserGroupsIdPrefix, id)
    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
        return conn.ExecCtx(ctx, query, id)
    }, userGroupsIdKey)
    return err
}

func (m *defaultUserGroupsModel) List(ctx context.Context, offset, limit int32) ([]*UserGroups, error) {
    query := fmt.Sprintf("SELECT %s FROM %s ORDER BY id DESC LIMIT ?, ?", userGroupsRows, m.table)

    var list []*UserGroups
    err := m.QueryRowsNoCacheCtx(ctx, &list, query, offset, limit)
    if err != nil {
        return nil, err
    }

    return list, nil
}

func (m *defaultUserGroupsModel) Count(ctx context.Context) (int64, error) {
    query := fmt.Sprintf("SELECT COUNT(*) FROM %s", m.table)

    var total int64
    err := m.QueryRowNoCacheCtx(ctx, &total, query)
    if err != nil {
        return 0, err
    }

    return total, nil
}
