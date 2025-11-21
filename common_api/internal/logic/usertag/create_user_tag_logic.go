package usertag

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type CreateUserTagLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewCreateUserTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserTagLogic {
    return &CreateUserTagLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *CreateUserTagLogic) CreateUserTag(req *types.CreateUserTagReq) (*types.CreateUserTagResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.CreateUserTag(l.ctx, &users.CreateUserTagReq{
        Name:        req.Name,
        DisplayName: req.DisplayName,
    })
    if err != nil {
        return nil, err
    }

    return &types.CreateUserTagResp{
        Id: rpcResp.Id,
    }, nil
}
