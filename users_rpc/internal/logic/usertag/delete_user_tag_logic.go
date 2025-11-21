package usertag

import (
    "context"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserTagLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewDeleteUserTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserTagLogic {
    return &DeleteUserTagLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *DeleteUserTagLogic) DeleteUserTag(in *users.DeleteUserTagReq) (*users.DeleteUserTagResp, error) {

    err := l.svcCtx.UserTagModel.Delete(l.ctx, in.Id)
    if err != nil {
        return nil, err
    }

    return &users.DeleteUserTagResp{
        Success: true,
    }, nil
}
