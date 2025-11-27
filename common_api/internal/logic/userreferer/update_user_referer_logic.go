package userreferer

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserRefererLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewUpdateUserRefererLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRefererLogic {
    return &UpdateUserRefererLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *UpdateUserRefererLogic) UpdateUserReferer(req *types.UpdateUserRefererReq) (*types.UpdateUserRefererResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.UpdateUserReferer(l.ctx, &users.UpdateUserRefererReq{
        Id:          req.Id,
        Name:        req.Name,
        DisplayName: req.DisplayName,
        ParentTree:  req.ParentTree,
    })
    if err != nil {
        return nil, err
    }

    return &types.UpdateUserRefererResp{
        Success: rpcResp.Success,
    }, nil
}
