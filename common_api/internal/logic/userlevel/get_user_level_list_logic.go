package userlevel

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserLevelListLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetUserLevelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLevelListLogic {
    return &GetUserLevelListLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *GetUserLevelListLogic) GetUserLevelList(req *types.GetUserLevelListReq) (*types.GetUserLevelListResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.GetUserLevelList(l.ctx, &users.GetUserLevelListReq{
        Page:     req.Page,
        PageSize: req.PageSize,
    })
    if err != nil {
        return nil, err
    }

    resp := &types.GetUserLevelListResp{
        Total: rpcResp.Total,
        List:  []types.UserLevelListItem{},
    }

    for _, item := range rpcResp.List {
        resp.List = append(resp.List, types.UserLevelListItem{
            Id:          item.Id,
            Name:        item.Name,
            DisplayName: item.DisplayName,
        })
    }

    return resp, nil
}
