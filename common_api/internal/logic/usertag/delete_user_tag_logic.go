package usertag

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type DeleteUserTagLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewDeleteUserTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserTagLogic {
    return &DeleteUserTagLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *DeleteUserTagLogic) DeleteUserTag(req *types.DeleteUserTagReq) (*types.DeleteUserTagResp, error) {

    _, err := l.svcCtx.UsersRpc.DeleteUserTag(l.ctx, &users.DeleteUserTagReq{
        Id: req.Id,
    })
    if err != nil {
        return nil, err
    }

    return &types.DeleteUserTagResp{
        Success: true,
    }, nil
}
