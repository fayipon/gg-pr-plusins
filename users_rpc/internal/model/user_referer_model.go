package model

import (
    "context"
    "database/sql"
    "fmt"

    "github.com/zeromicro/go-zero/core/stores/cache"
    "github.com/zeromicro/go-zero/core/stores/sqlc"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserRefererModel = (*defaultUserRefererModel)(nil)

// ======================================================
// cache prefix
// ======================================================
const (
    cacheUserRefererIdPrefix = "cache:userReferer:id:"

    userRefererRows = "`id`,`user_id`,`parent_tree`,`name`,`display_name`,`visit_count`,`register_count`,`first_deposit_count`,`created_at`,`updated_at`"

    userRefererRowsExpectAutoSet = "`user_id`,`parent_tree`,`name`,`display_name`,`visit_count`,`register_count`,`first_deposit_count`,`created_at`,`updated_at`"

    userRefererRowsWithPlaceHolder = "`user_id`=?,`parent_tree`=?,`name`=?,`display_name`=?,`visit_count`=?,`register_count`=?,`first_deposit_count`=?,`created_at`=?,`updated_at`=?"
)

// ======================================================
// 数据结构
// ======================================================
type (
    UserReferer struct {
        Id                uint64 `db:"id"`
        UserId            uint64 `db:"user_id"`
        ParentTree        string `db:"parent_tree"`
        Name              string `db:"name"`
        DisplayName       string `db:"display_name"`
        VisitCount        int64  `db:"visit_count"`
        RegisterCount     int64  `db:"register_count"`
        FirstDepositCount int64  `db:"first_deposit_count"`
        CreatedAt         int64  `db:"created_at"`
        UpdatedAt         int64  `db:"updated_at"`
    }

    UserRefererModel interface {
        Insert(ctx context.Context, data *UserReferer) (sql.Result, error)
        FindOne(ctx context.Context, id uint64) (*UserReferer, error)
        FindByUserId(ctx context.Context, userId uint64) (*UserReferer, error)
        Update(ctx context.Context, data *UserReferer) error
        Delete(ctx context.Context, id uint64) error

        List(ctx context.Context, offset, limit int32) ([]*UserReferer, error)
        Count(ctx context.Context) (int64, error)

		FindOneByUserId(ctx context.Context, userId uint64) (*UserReferer, error)
    }

    defaultUserRefererModel struct {
        sqlc.CachedConn
        table string
    }
)

// ======================================================
// Model 构造函数
// ======================================================
func NewUserRefererModel(conn sqlx.SqlConn, c cache.CacheConf) UserRefererModel {
    return &defaultUserRefererModel{
        CachedConn: sqlc.NewConn(conn, c),
        table:      "`user_referer`",
    }
}

// ======================================================
// Insert（带缓存）
// ======================================================
func (m *defaultUserRefererModel) Insert(ctx context.Context, data *UserReferer) (sql.Result, error) {
    userRefererIdKey := fmt.Sprintf("%s%v", cacheUserRefererIdPrefix, data.Id)

    ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        query := fmt.Sprintf(
            "INSERT INTO %s (%s) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
            m.table,
            userRefererRowsExpectAutoSet,
        )

        return conn.ExecCtx(ctx, query,
            data.UserId,
            data.ParentTree,
            data.Name,
            data.DisplayName,
            data.VisitCount,
            data.RegisterCount,
            data.FirstDepositCount,
            data.CreatedAt,
            data.UpdatedAt,
        )
    }, userRefererIdKey)

    return ret, err
}

// ======================================================
// FindOne（带缓存）
// ======================================================
func (m *defaultUserRefererModel) FindOne(ctx context.Context, id uint64) (*UserReferer, error) {
    userRefererIdKey := fmt.Sprintf("%s%v", cacheUserRefererIdPrefix, id)

    var resp UserReferer
    err := m.QueryRowCtx(ctx, &resp, userRefererIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
        query := fmt.Sprintf("SELECT %s FROM %s WHERE `id` = ? LIMIT 1", userRefererRows, m.table)
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

// ======================================================
// FindByUserId（无缓存）
// ======================================================
func (m *defaultUserRefererModel) FindByUserId(ctx context.Context, userId uint64) (*UserReferer, error) {
    query := fmt.Sprintf("SELECT %s FROM %s WHERE `user_id` = ? LIMIT 1", userRefererRows, m.table)

    var resp UserReferer
    err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId)

    switch err {
    case nil:
        return &resp, nil
    case sqlc.ErrNotFound:
        return nil, ErrNotFound
    default:
        return nil, err
    }
}

// ======================================================
// Update（带缓存）
// ======================================================
func (m *defaultUserRefererModel) Update(ctx context.Context, data *UserReferer) error {
    userRefererIdKey := fmt.Sprintf("%s%v", cacheUserRefererIdPrefix, data.Id)

    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        query := fmt.Sprintf(
            "UPDATE %s SET %s WHERE `id` = ?",
            m.table,
            userRefererRowsWithPlaceHolder,
        )

        return conn.ExecCtx(ctx, query,
            data.UserId,
            data.ParentTree,
            data.Name,
            data.DisplayName,
            data.VisitCount,
            data.RegisterCount,
            data.FirstDepositCount,
            data.CreatedAt,
            data.UpdatedAt,
            data.Id,
        )
    }, userRefererIdKey)

    return err
}

// ======================================================
// Delete（带缓存）
// ======================================================
func (m *defaultUserRefererModel) Delete(ctx context.Context, id uint64) error {
    userRefererIdKey := fmt.Sprintf("%s%v", cacheUserRefererIdPrefix, id)

    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        query := fmt.Sprintf("DELETE FROM %s WHERE `id` = ?", m.table)
        return conn.ExecCtx(ctx, query, id)
    }, userRefererIdKey)

    return err
}

// ======================================================
// List（无缓存）
// ======================================================
func (m *defaultUserRefererModel) List(ctx context.Context, offset, limit int32) ([]*UserReferer, error) {
    query := fmt.Sprintf(
        "SELECT %s FROM %s ORDER BY id DESC LIMIT ?, ?",
        userRefererRows,
        m.table,
    )

    var list []*UserReferer
    err := m.QueryRowsNoCacheCtx(ctx, &list, query, offset, limit)

    return list, err
}

// ======================================================
// Count（无缓存）
// ======================================================
func (m *defaultUserRefererModel) Count(ctx context.Context) (int64, error) {
    query := fmt.Sprintf("SELECT COUNT(*) FROM %s", m.table)

    var total int64
    err := m.QueryRowNoCacheCtx(ctx, &total, query)

    return total, err
}

// =======================
// FindOneByUserId （不缓存）
// =======================
func (m *defaultUserRefererModel) FindOneByUserId(ctx context.Context, userId uint64) (*UserReferer, error) {

    query := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = ? LIMIT 1",
        userRefererRows,
        m.table,
    )

    var resp UserReferer
    err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId)
    switch err {
    case nil:
        return &resp, nil
    case sqlc.ErrNotFound:
        return nil, ErrNotFound
    default:
        return nil, err
    }
}