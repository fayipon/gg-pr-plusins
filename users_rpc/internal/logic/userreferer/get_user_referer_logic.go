package userreferer

import (
    "context"

    "github.com/zeromicro/go-zero/core/logx"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
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

func (l *GetUserRefererLogic) GetUserReferer(req *users.GetUserRefererReq) (*users.GetUserRefererResp, error) {

    data, err := l.svcCtx.UserRefererModel.FindOne(l.ctx, req.Id)
    if err != nil {
        return nil, err
    }

    return &users.GetUserRefererResp{
        Info: &users.UserRefererInfo{
            Id:                data.Id,
            UserId:            data.UserId,
            ParentTree:        data.ParentTree,
            Name:              data.Name,
            DisplayName:       data.DisplayName,
            VisitCount:        int64(data.VisitCount),
            RegisterCount:     int64(data.RegisterCount),
            FirstDepositCount: int64(data.FirstDepositCount),
            CreatedAt:         data.CreatedAt,
            UpdatedAt:         data.UpdatedAt,
        },
    }, nil
}
