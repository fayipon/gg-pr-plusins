package usergroup

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"

    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserGroupLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewDeleteUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserGroupLogic {
    return &DeleteUserGroupLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *DeleteUserGroupLogic) DeleteUserGroup(req *types.DeleteUserGroupReq) (*types.DeleteUserGroupResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.DeleteUserGroup(l.ctx, &users.DeleteUserGroupReq{
        Id: req.Id,
    })

    if err != nil {
        return nil, err
    }

    return &types.DeleteUserGroupResp{
        Success: rpcResp.Success,
    }, nil
}
