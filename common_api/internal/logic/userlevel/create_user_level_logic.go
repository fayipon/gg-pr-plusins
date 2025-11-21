package userlevel

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type CreateUserLevelLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewCreateUserLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLevelLogic {
    return &CreateUserLevelLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *CreateUserLevelLogic) CreateUserLevel(req *types.CreateUserLevelReq) (*types.CreateUserLevelResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.CreateUserLevel(l.ctx, &users.CreateUserLevelReq{
        Name:        req.Name,
        DisplayName: req.DisplayName,
        Setting:     req.Setting,
    })
    if err != nil {
        return nil, err
    }

    return &types.CreateUserLevelResp{
        Id: rpcResp.Id,
    }, nil
}
