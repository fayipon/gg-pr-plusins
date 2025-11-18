package logic

import (
    "context"

    "users_rpc/internal/svc"
    "users_rpc/users"

    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
    return &DeleteUserLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *DeleteUserLogic) DeleteUser(in *users.DeleteUserReq) (*users.DeleteUserResp, error) {

    err := l.svcCtx.UsersModel.Delete(l.ctx, in.Id)
    if err != nil {
        r
