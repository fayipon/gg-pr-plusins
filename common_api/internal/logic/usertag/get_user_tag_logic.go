package usertag

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type GetUserTagLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetUserTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTagLogic {
    return &GetUserTagLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetUserTagLogic) GetUserTag(req *types.GetUserTagReq) (*types.GetUserTagResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.GetUserTag(l.ctx, &users.GetUserTagReq{
        Id: req.Id,
    })
	
    if err != nil {
        return nil, err
    }

    return &types.GetUserTagResp{
        Id:          rpcResp.Id,
        Name:        rpcResp.Name,
        DisplayName: rpcResp.DisplayName,
        CreatedAt:   rpcResp.CreatedAt,
        UpdatedAt:   rpcResp.UpdatedAt,
    }, nil
}
