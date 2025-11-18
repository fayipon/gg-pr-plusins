package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Users struct {
	Id        int64     `db:"id"`
	Account   string    `db:"account"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type UsersModel interface {
	Insert(ctx context.Context, data *Users) (int64, error)
	FindOne(ctx context.Context, id int64) (*Users, error)
	Update(ctx context.Context, data *Users) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]Users, error)
}

type defaultUsersModel struct {
	conn sqlx.SqlConn
}

func NewUsersModel(conn sqlx.SqlConn) UsersModel {
	return &defaultUsersModel{conn: conn}
}
