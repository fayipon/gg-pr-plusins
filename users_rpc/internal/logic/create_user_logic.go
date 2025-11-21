package logic

import (
    "context"
    "time"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/model"
    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
    return &CreateUserLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *CreateUserLogic) CreateUser(in *users.CreateUserReq) (*users.CreateUserResp, error) {

    now := time.Now().Unix() 

    u := &model.Users{
        Account:  in.Account,
        Password: in.Password,
        Status:   1,
        CreatedAt: now,
        UpdatedAt: now,
    }

    _, err := l.svcCtx.UserModel.Insert(l.ctx, u)
    if err != nil {
        return nil, err
    }

    return &users.CreateUserResp{
        Success: true,
    }, nil
}
