package userlevel

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserLevelListLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetUserLevelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLevelListLogic {
    return &GetUserLevelListLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetUserLevelListLogic) GetUserLevelList(req *types.GetUserLevelListReq) (*types.GetUserLevelListResp, error) {

    rpcReq := &users.GetUserLevelListReq{
        Page:     req.Page,
        PageSize: req.PageSize,
    }

    rpcResp, err := l.svcCtx.UsersRpc.GetUserLevelList(l.ctx, rpcReq)
    if err != nil {
        return nil, err
    }

    resp := &types.GetUserLevelListResp{
        Total: rpcResp.Total,
        List:  make([]*types.UserLevelListItem, 0),
    }

    for _, g := range rpcResp.List {
        resp.List = append(resp.List, &types.UserLevelListItem{
            Id:          g.Id,
            Name:        g.Name,
            DisplayName: g.DisplayName,
        })
    }

    return resp, nil
}
