package logic

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
	"github.com/fayipon/gg-pr-plusins/users_rpc/users"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *users.CreateUserReq) (*users.CreateUserResp, error) {

	qb := squirrel.Insert("users").
		Columns("account", "password", "status", "level_id", "created_at").
		Values(req.Account, req.Password, 1, 1, time.Now())

	sqlStr, args, _ := qb.ToSql()
	_, err := l.svcCtx.DB.ExecCtx(l.ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}

	return &users.CreateUserResp{
		Success: true,
	}, nil
}
