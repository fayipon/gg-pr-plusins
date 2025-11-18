package logic

import (
    "context"

    "users_rpc/internal/svc"
    "users_rpc/internal/model"
    "users_rpc/users"

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

func (l *CreateUserLogic) CreateUser(in *users.CreateUserReq) (*users.CreateUserResp, error) {

    data := &model.Users{
        Account:  in.Account,
        Password: in.Password,
    }

    id, err := l.svcCtx.UsersModel.Insert(l.ctx, data)
    if err != nil {
        return nil, err
    }

    return &users.CreateUserResp{
        Id: id,
    }, nil
}
