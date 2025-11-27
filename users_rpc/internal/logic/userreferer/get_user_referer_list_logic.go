package userreferer

import (
    "context"

    "github.com/zeromicro/go-zero/core/logx"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
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

func (l *GetUserRefererListLogic) GetUserRefererList(req *users.GetUserRefererListReq) (*users.GetUserRefererListResp, error) {

    offset := (req.Page - 1) * req.PageSize
    list, err := l.svcCtx.UserRefererModel.List(l.ctx, offset, req.PageSize)
    if err != nil {
        return nil, err
    }

    count, err := l.svcCtx.UserRefererModel.Count(l.ctx)
    if err != nil {
        return nil, err
    }

    respList := make([]*users.UserRefererInfo, 0)
    for _, v := range list {
        respList = append(respList, &users.UserRefererInfo{
            Id:                v.Id,
            UserId:            v.UserId,
            ParentTree:        v.ParentTree,
            Name:              v.Name,
            DisplayName:       v.DisplayName,
            VisitCount:        int64(v.VisitCount),
            RegisterCount:     int64(v.RegisterCount),
            FirstDepositCount: int64(v.FirstDepositCount),
            CreatedAt:         v.CreatedAt,
            UpdatedAt:         v.UpdatedAt,
        })
    }

    return &users.GetUserRefererListResp{
        Total: count,
        List:  respList,
    }, nil
}
