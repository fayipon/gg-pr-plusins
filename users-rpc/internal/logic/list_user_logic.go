package logic

import (
    "context"

    "users-rpc/internal/svc"
    "users-rpc/users"

    "github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
    return &ListUserLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *ListUserLogic) ListUser(in *users.ListUserReq) (*users.ListUserResp, error) {

    list, err := l.svcCtx.UsersModel.List(l.ctx)
    if err != nil {
        return nil, err
    }

    resp := &users.ListUserResp{}

    for _, u := range list {
        resp.List = append(resp.List, &users.UserInfo{
            Id:       u.Id,
            Account:  u.Account,
            Password: u.Password,
        })
    }

    return resp, nil
}
