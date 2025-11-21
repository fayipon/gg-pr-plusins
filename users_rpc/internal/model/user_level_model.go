package model

import (
    "context"
    "database/sql"
    "fmt"
    "strings"

    "github.com/zeromicro/go-zero/core/stores/cache"
    "github.com/zeromicro/go-zero/core/stores/sqlc"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserLevelsModel = (*defaultUserLevelsModel)(nil)

const (
    cacheUserLevelsIdPrefix = "cache:userLevels:id:"

    userLevelsRows = "`id`,`name`,`display_name`,`setting`,`created_at`,`updated_at`"

    userLevelsRowsExpectAutoSet = "`name`,`display_name`,`setting`,`created_at`,`updated_at`"

    userLevelsRowsWithPlaceHolder = "`name`=?,`display_name`=?,`setting`=?,`created_at`=?,`updated_at`=?"
)

type (
    UserLevels struct {
        Id          uint64 `db:"id"`
        Name        string `db:"name"`
        DisplayName string `db:"display_name"`
        Setting     string `db:"setting"`
        CreatedAt   int64  `db:"created_at"`
        UpdatedAt   int64  `db:"updated_at"`
    }

    UserLevelsModel interface {
        Insert(ctx context.Context, data *UserLevels) (sql.Result, error)
        FindOne(ctx context.Context, id uint64) (*UserLevels, error)
        Update(ctx context.Context, data *UserLevels) error
        Delete(ctx context.Context, id uint64) error
        FindByIds(ctx context.Context, ids []uint64) ([]*UserLevels, error)
        
        List(ctx context.Context, page, pageSize int32) ([]*UserLevels, error)
        Count(ctx context.Context) (int64, error)
    }

    defaultUserLevelsModel struct {
        sqlc.CachedConn
        table string
    }
)

// =====================
// Model 构造函数
// =====================
func NewUserLevelsModel(conn sqlx.SqlConn, c cache.CacheConf) UserLevelsModel {
    return &defaultUserLevelsModel{
        CachedConn: sqlc.NewConn(conn, c),
        table:      "`user_levels`",
    }
}

// =====================
// Insert
// =====================
func (m *defaultUserLevelsModel) Insert(ctx context.Context, data *UserLevels) (sql.Result, error) {
    userLevelsIdKey := fmt.Sprintf("%s%v", cacheUserLevelsIdPrefix, data.Id)
    ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userLevelsRowsExpectAutoSet)
        return conn.ExecCtx(ctx, query, data.Name, data.DisplayName, data.Setting, data.CreatedAt, data.UpdatedAt)
    }, userLevelsIdKey)
    return ret, err
}

// =====================
// FindOne（带缓存）
// =====================
func (m *defaultUserLevelsModel) FindOne(ctx context.Context, id uint64) (*UserLevels, error) {
    userLevelsIdKey := fmt.Sprintf("%s%v", cacheUserLevelsIdPrefix, id)
    var resp UserLevels
    err := m.QueryRowCtx(ctx, &resp, userLevelsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
        query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userLevelsRows, m.table)
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

// =====================
// Update（带缓存）
// =====================
func (m *defaultUserLevelsModel) Update(ctx context.Context, data *UserLevels) error {
    userLevelsIdKey := fmt.Sprintf("%s%v", cacheUserLevelsIdPrefix, data.Id)
    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userLevelsRowsWithPlaceHolder)
        return conn.ExecCtx(ctx, query, data.Name, data.DisplayName, data.Setting, data.CreatedAt, data.UpdatedAt, data.Id)
    }, userLevelsIdKey)
    return err
}

// =====================
// Delete（带缓存）
// =====================
func (m *defaultUserLevelsModel) Delete(ctx context.Context, id uint64) error {
    userLevelsIdKey := fmt.Sprintf("%s%v", cacheUserLevelsIdPrefix, id)

    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
        return conn.ExecCtx(ctx, query, id)
    }, userLevelsIdKey)

    return err
}

// =====================
// ⭐ FindByIds（批量查询，无缓存）
// =====================
func (m *defaultUserLevelsModel) FindByIds(ctx context.Context, ids []uint64) ([]*UserLevels, error) {
    if len(ids) == 0 {
        return []*UserLevels{}, nil
    }

    placeholders := make([]string, len(ids))
    args := make([]any, len(ids))

    for i, id := range ids {
        placeholders[i] = "?"
        args[i] = id
    }

    query := fmt.Sprintf(
        "SELECT %s FROM %s WHERE id IN (%s)",
        userLevelsRows,
        m.table,
        strings.Join(placeholders, ","),
    )

    var resp []*UserLevels
    err := m.QueryRowsNoCacheCtx(ctx, &resp, query, args...)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

func (m *defaultUserLevelsModel) List(ctx context.Context, offset, limit int32) ([]*UserLevels, error) {
    query := fmt.Sprintf(
        "SELECT %s FROM %s ORDER BY id DESC LIMIT ?, ?",
        userLevelsRows,
        m.table,
    )

    var list []*UserLevels
    err := m.QueryRowsNoCacheCtx(ctx, &list, query, offset, limit)
    if err != nil {
        return nil, err
    }

    return list, nil
}


func (m *defaultUserLevelsModel) Count(ctx context.Context) (int64, error) {
    query := fmt.Sprintf("SELECT COUNT(*) FROM %s", m.table)

    var total int64
    err := m.QueryRowNoCacheCtx(ctx, &total, query)
    if err != nil {
        return 0, err
    }

    return total, nil
}

