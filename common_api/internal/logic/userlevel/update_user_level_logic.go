package userlevel

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLevelLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewUpdateUserLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLevelLogic {
    return &UpdateUserLevelLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *UpdateUserLevelLogic) UpdateUserLevel(req *types.UpdateUserLevelReq) (*types.UpdateUserLevelResp, error) {

    _, err := l.svcCtx.UsersRpc.UpdateUserLevel(l.ctx, &users.UpdateUserLevelReq{
        Id:          req.Id,
        Name:        req.Name,
        DisplayName: req.DisplayName,
        Setting:     req.Setting,
    })
    if err != nil {
        return nil, err
    }

    return &types.UpdateUserLevelResp{
        Success: true,
    }, nil
}
