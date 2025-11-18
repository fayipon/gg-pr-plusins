package model

import (
	"database/sql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Users struct {
	Id       int64  `db:"id"`
	Account  string `db:"account"`
	Password string `db:"password"`
}

type UsersModel interface {
	FindOne(id int64) (*Users, error)
	FindByAccount(account string) (*Users, error)
	Insert(data *Users) (int64, error)
	List(offset, limit int32, keyword string) ([]*Users, error)
	Count(keyword string) (int64, error)
}

type customUsersModel struct {
	conn sqlx.SqlConn
}

func NewUsersModel(datasource string) UsersModel {
	return &customUsersModel{
		conn: sqlx.NewMysql(datasource),
	}
}

func (m *customUsersModel) FindOne(id int64) (*Users, error) {
	query := `SELECT id, account, password FROM users WHERE id = ? LIMIT 1`
	var resp Users
	err := m.conn.QueryRow(&resp, query, id)

	switch err {
	case nil:
		return &resp, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customUsersModel) FindByAccount(account string) (*Users, error) {
	query := `SELECT id, account, password FROM users WHERE account = ? LIMIT 1`
	var resp Users
	err := m.conn.QueryRow(&resp, query, account)

	switch err {
	case nil:
		return &resp, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customUsersModel) Insert(data *Users) (int64, error) {
	query := `
	INSERT INTO users (account, password, created_at, updated_at)
	VALUES (?, ?, NOW(), NOW())`

	result, err := m.conn.Exec(query, data.Account, data.Password)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (m *customUsersModel) List(offset, limit int32, keyword string) ([]*Users, error) {
	query := `SELECT id, account, password FROM users WHERE account LIKE ? ORDER BY id DESC LIMIT ? OFFSET ?`
	like := "%" + keyword + "%"

	var resp []*Users
	err := m.conn.QueryRows(&resp, query, like, limit, offset)
	return resp, err
}

func (m *customUsersModel) Count(keyword string) (int64, error) {
	query := `SELECT COUNT(*) as cnt FROM users WHERE account LIKE ?`
	like := "%" + keyword + "%"

	var cnt int64
	err := m.conn.QueryRow(&cnt, query, like)
	return cnt, err
}