package usertag

import (
    "context"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserTagLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetUserTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTagLogic {
    return &GetUserTagLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *GetUserTagLogic) GetUserTag(in *users.GetUserTagReq) (*users.GetUserTagResp, error) {

    tag, err := l.svcCtx.UserTagModel.FindOne(l.ctx, in.Id)
    if err != nil {
        return nil, err
    }

    return &users.GetUserTagResp{
        Id:          tag.Id,
        Name:        tag.Name,
        DisplayName: tag.DisplayName,
        CreatedAt:   tag.CreatedAt,
        UpdatedAt:   tag.UpdatedAt,
    }, nil
}
