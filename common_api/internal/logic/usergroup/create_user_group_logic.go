package usergroup

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"

    "github.com/zeromicro/go-zero/core/logx"
)

type CreateUserGroupLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewCreateUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserGroupLogic {
    return &CreateUserGroupLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *CreateUserGroupLogic) CreateUserGroup(req *types.CreateUserGroupReq) (*types.CreateUserGroupResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.CreateUserGroup(l.ctx, &users.CreateUserGroupReq{
        Name:        req.Name,
        DisplayName: req.DisplayName,
        Setting:     req.Setting,
    })

    if err != nil {
        return nil, err
    }

    return &types.CreateUserGroupResp{
        Id: rpcResp.Id,
    }, nil
}
