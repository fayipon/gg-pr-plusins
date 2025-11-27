package userreferer

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserRefererLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewDeleteUserRefererLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserRefererLogic {
    return &DeleteUserRefererLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *DeleteUserRefererLogic) DeleteUserReferer(req *types.DeleteUserRefererReq) (*types.DeleteUserRefererResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.DeleteUserReferer(l.ctx, &users.DeleteUserRefererReq{
        Id: req.Id,
    })
    if err != nil {
        return nil, err
    }

    return &types.DeleteUserRefererResp{
        Success: rpcResp.Success,
    }, nil
}
