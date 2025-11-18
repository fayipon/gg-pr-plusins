package logic

import (
    "context"

    "users_rpc/internal/svc"
    "users_rpc/users"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
    return &GetUserLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetUserLogic) GetUser(in *users.GetUserReq) (*users.GetUserResp, error) {

    u, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.Id)
    if err != nil {
        return nil, err
    }

    return &users.GetUserResp{
        User: &users.UserInfo{
            Id:       u.Id,
            Account:  u.Account,
            Password: u.Password,
        },
    }, nil
}
