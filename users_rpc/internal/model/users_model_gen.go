// Code generated manually to match go-zero style.
// DO NOT EDIT unless you know the model structure.

package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	usersTable          = "`users`"
	cacheUsersIdPrefix  = "cache:users:id:"
)

type Users struct {
	Id        int64     `db:"id"`
	Account   string    `db:"account"`
	Password  string    `db:"password"`
	CreatedAt string    `db:"created_at"`
	UpdatedAt string    `db:"updated_at"`
}

type usersModel interface {
	Insert(ctx context.Context, data *Users) (int64, error)
	FindOne(ctx context.Context, id int64) (*Users, error)
	Update(ctx context.Context, data *Users) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*Users, error)
}

type defaultUsersModel struct {
	sqlc.CachedConn
	table string
}

// newUsersModel 创建自动生成模型
func newUsersModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewCachedConn(conn, c),
		table:      usersTable,
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (int64, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (account, password, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id",
		m.table,
	)

	var id int64
	err := m.QueryRowNoCacheCtx(ctx, &id, query, data.Account, data.Password)
	return id, err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	usersKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)

	var resp Users
	query := fmt.Sprintf(
		"SELECT id, account, password, created_at, updated_at FROM %s WHERE id = $1",
		m.table,
	)

	err := m.QueryRowCtx(ctx, &resp, usersKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		return conn.QueryRowCtx(ctx, v, query, id)
	})

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, sqlc.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Update(ctx context.Context, data *Users) error {
	usersKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)

	query := fmt.Sprintf(
		"UPDATE %s SET account=$1, password=$2, updated_at=NOW() WHERE id=$3",
		m.table,
	)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, data.Account, data.Password, data.Id)
	}, usersKey)

	return err
}

func (m *defaultUsersModel) Delete(ctx context.Context, id int64) error {
	usersKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", m.table)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, id)
	}, usersKey)

	return err
}

func (m *defaultUsersModel) List(ctx context.Context) ([]*Users, error) {

	query := fmt.Sprintf("SELECT id, account, password, created_at, updated_at FROM %s ORDER BY id", m.table)

	var resp []*Users
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
