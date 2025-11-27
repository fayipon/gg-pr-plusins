package userreferer

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type CreateUserRefererLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewCreateUserRefererLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserRefererLogic {
    return &CreateUserRefererLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *CreateUserRefererLogic) CreateUserReferer(req *types.CreateUserRefererReq) (*types.CreateUserRefererResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.CreateUserReferer(l.ctx, &users.CreateUserRefererReq{
        UserId:      req.UserId,
        ParentTree:  req.ParentTree,
        Name:        req.Name,
        DisplayName: req.DisplayName,
    })
    if err != nil {
        return nil, err
    }

    return &types.CreateUserRefererResp{
        Id: rpcResp.Id,
    }, nil
}
