package logic

import (
    "context"

    "users-rpc/internal/svc"
    "users-rpc/internal/model"
    "users-rpc/users"

    "github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
    return &UpdateUserLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *UpdateUserLogic) UpdateUser(in *users.UpdateUserReq) (*users.UpdateUserResp, error) {

    data := &model.Users{
        Id:       in.Id,
        Account:  in.Account,
        Password: in.Password,
    }

    err := l.svcCtx.UsersModel.Update(l.ctx, data)
    if err != nil {
        return nil, err
    }

    return &users.UpdateUserResp{
        Success: true,
    }, nil
}
