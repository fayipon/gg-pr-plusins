package userreferer

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserRefererListLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetUserRefererListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRefererListLogic {
    return &GetUserRefererListLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetUserRefererListLogic) GetUserRefererList(req *types.GetUserRefererListReq) (*types.GetUserRefererListResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.GetUserRefererList(l.ctx, &users.GetUserRefererListReq{
        Page:     req.Page,
        PageSize: req.PageSize,
    })
    if err != nil {
        return nil, err
    }

    list := make([]types.UserRefererInfo, 0, len(rpcResp.List))
    for _, item := range rpcResp.List {
        list = append(list, types.UserRefererInfo{
            Id:                item.Id,
            UserId:            item.UserId,
            ParentTree:        item.ParentTree,
            Name:              item.Name,
            DisplayName:       item.DisplayName,
            VisitCount:        item.VisitCount,
            RegisterCount:     item.RegisterCount,
            FirstDepositCount: item.FirstDepositCount,
            CreatedAt:         item.CreatedAt,
            UpdatedAt:         item.UpdatedAt,
        })
    }

    return &types.GetUserRefererListResp{
        Total: rpcResp.Total,
        List:  list,
    }, nil
}
