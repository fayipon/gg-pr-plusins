package userlevel

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
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

func (l *DeleteUserLevelLogic) DeleteUserLevel(req *types.DeleteUserLevelReq) (*types.DeleteUserLevelResp, error) {

    _, err := l.svcCtx.UsersRpc.DeleteUserLevel(l.ctx, &users.DeleteUserLevelReq{
        Id: req.Id,
    })
    if err != nil {
        return nil, err
    }

    return &types.DeleteUserLevelResp{
        Success: true,
    }, nil
}
