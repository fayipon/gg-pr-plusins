package userreferer

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserRefererLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetUserRefererLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRefererLogic {
    return &GetUserRefererLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetUserRefererLogic) GetUserReferer(req *types.GetUserRefererReq) (*types.GetUserRefererResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.GetUserReferer(l.ctx, &users.GetUserRefererReq{
        Id: req.Id,
    })
    if err != nil {
        return nil, err
    }

    info := rpcResp.Info

    return &types.GetUserRefererResp{
        Info: types.UserRefererInfo{
            Id:                info.Id,
            UserId:            info.UserId,
            ParentTree:        info.ParentTree,
            Name:              info.Name,
            DisplayName:       info.DisplayName,
            VisitCount:        info.VisitCount,
            RegisterCount:     info.RegisterCount,
            FirstDepositCount: info.FirstDepositCount,
            CreatedAt:         info.CreatedAt,
            UpdatedAt:         info.UpdatedAt,
        },
    }, nil
}
