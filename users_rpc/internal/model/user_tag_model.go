package model

import (
    "context"
    "database/sql"
    "fmt"

    "github.com/zeromicro/go-zero/core/stores/cache"
    "github.com/zeromicro/go-zero/core/stores/sqlc"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserTagModel = (*defaultUserTagModel)(nil)

const (
    cacheUserTagIdPrefix = "cache:userTag:id:"

    userTagRows = "`id`,`name`,`display_name`,`created_at`,`updated_at`"

    userTagRowsExpectAutoSet = "`name`,`display_name`,`created_at`,`updated_at`"

    userTagRowsWithPlaceHolder = "`name`=?,`display_name`=?,`created_at`=?,`updated_at`=?"
)

type (
    UserTag struct {
        Id          uint64 `db:"id"`
        Name        string `db:"name"`
        DisplayName string `db:"display_name"`
        CreatedAt   int64  `db:"created_at"`
        UpdatedAt   int64  `db:"updated_at"`
    }

    UserTagModel interface {
        Insert(ctx context.Context, data *UserTag) (sql.Result, error)
        FindOne(ctx context.Context, id uint64) (*UserTag, error)
        Update(ctx context.Context, data *UserTag) error
        Delete(ctx context.Context, id uint64) error

        List(ctx context.Context, offset, limit int32) ([]*UserTag, error)
        Count(ctx context.Context) (int64, error)
    }

    defaultUserTagModel struct {
        sqlc.CachedConn
        table string
    }
)

func NewUserTagModel(conn sqlx.SqlConn, c cache.CacheConf) UserTagModel {
    return &defaultUserTagModel{
        CachedConn: sqlc.NewConn(conn, c),
        table:      "`user_tags`",
    }
}

func (m *defaultUserTagModel) Insert(ctx context.Context, data *UserTag) (sql.Result, error) {
    userTagIdKey := fmt.Sprintf("%s%v", cacheUserTagIdPrefix, data.Id)
    ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userTagRowsExpectAutoSet)
        return conn.ExecCtx(ctx, query, data.Name, data.DisplayName, data.CreatedAt, data.UpdatedAt)
    }, userTagIdKey)
    return ret, err
}

func (m *defaultUserTagModel) FindOne(ctx context.Context, id uint64) (*UserTag, error) {
    userTagIdKey := fmt.Sprintf("%s%v", cacheUserTagIdPrefix, id)
    var resp UserTag
    err := m.QueryRowCtx(ctx, &resp, userTagIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
        query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userTagRows, m.table)
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

func (m *defaultUserTagModel) Update(ctx context.Context, data *UserTag) error {
    userTagIdKey := fmt.Sprintf("%s%v", cacheUserTagIdPrefix, data.Id)
    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userTagRowsWithPlaceHolder)
        return conn.ExecCtx(ctx, query, data.Name, data.DisplayName, data.CreatedAt, data.UpdatedAt, data.Id)
    }, userTagIdKey)
    return err
}

func (m *defaultUserTagModel) Delete(ctx context.Context, id uint64) error {
    userTagIdKey := fmt.Sprintf("%s%v", cacheUserTagIdPrefix, id)
    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
        return conn.ExecCtx(ctx, query, id)
    }, userTagIdKey)
    return err
}

func (m *defaultUserTagModel) List(ctx context.Context, offset, limit int32) ([]*UserTag, error) {
    query := fmt.Sprintf(
        "SELECT %s FROM %s ORDER BY id DESC LIMIT ?, ?",
        userTagRows,
        m.table,
    )

    var list []*UserTag
    err := m.QueryRowsNoCacheCtx(ctx, &list, query, offset, limit)
    if err != nil {
        return nil, err
    }

    return list, nil
}

func (m *defaultUserTagModel) Count(ctx context.Context) (int64, error) {
    query := fmt.Sprintf("SELECT COUNT(*) FROM %s", m.table)

    var total int64
    err := m.QueryRowNoCacheCtx(ctx, &total, query)
    if err != nil {
        return 0, err
    }

    return total, nil
}
