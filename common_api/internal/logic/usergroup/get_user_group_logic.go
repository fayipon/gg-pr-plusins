package usergroup

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserGroupLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserGroupLogic {
    return &GetUserGroupLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *GetUserGroupLogic) GetUserGroup(req *types.GetUserGroupReq) (*types.GetUserGroupResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.GetUserGroup(l.ctx, &users.GetUserGroupReq{
        Id: req.Id,
    })

    if err != nil {
        return nil, err
    }

    return &types.GetUserGroupResp{
        Id:          rpcResp.Id,
        Name:        rpcResp.Name,
        DisplayName: rpcResp.DisplayName,
        Setting:     rpcResp.Setting,
        CreatedAt:   rpcResp.CreatedAt,
        UpdatedAt:   rpcResp.UpdatedAt,
    }, nil
}
