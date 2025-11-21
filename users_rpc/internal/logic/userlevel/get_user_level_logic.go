package userlevel

import (
    "context"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserLevelLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetUserLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLevelLogic {
    return &GetUserLevelLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *GetUserLevelLogic) GetUserLevel(in *users.GetUserLevelReq) (*users.GetUserLevelResp, error) {

    lv, err := l.svcCtx.UserLevelModel.FindOne(l.ctx, in.Id)
    if err != nil {
        return nil, err
    }

    return &users.GetUserLevelResp{
        Id:         lv.Id,
        Name:       lv.Name,
        DisplayName: lv.DisplayName,
        Setting:    lv.Setting,
        CreatedAt:  lv.CreatedAt,
        UpdatedAt:  lv.UpdatedAt,
    }, nil
}
