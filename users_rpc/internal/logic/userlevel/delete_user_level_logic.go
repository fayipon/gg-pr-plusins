package userlevel

import (
    "context"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLevelLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewDeleteUserLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLevelLogic {
    return &DeleteUserLevelLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *DeleteUserLevelLogic) DeleteUserLevel(in *users.DeleteUserLevelReq) (*users.DeleteUserLevelResp, error) {

    err := l.svcCtx.UserLevelModel.Delete(l.ctx, in.Id)
    if err != nil {
        return nil, err
    }

    return &users.DeleteUserLevelResp{
        Success: true,
    }, nil
}
