package usergroup

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"

    "github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserGroupLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewUpdateUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserGroupLogic {
    return &UpdateUserGroupLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *UpdateUserGroupLogic) UpdateUserGroup(req *types.UpdateUserGroupReq) (*types.UpdateUserGroupResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.UpdateUserGroup(l.ctx, &users.UpdateUserGroupReq{
        Id:          req.Id,
        Name:        req.Name,
        DisplayName: req.DisplayName,
    })

    if err != nil {
        return nil, err
    }

    return &types.UpdateUserGroupResp{
        Success: rpcResp.Success,
    }, nil
}
