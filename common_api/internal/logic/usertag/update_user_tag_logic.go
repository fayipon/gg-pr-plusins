package usertag

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type UpdateUserTagLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewUpdateUserTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserTagLogic {
    return &UpdateUserTagLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *UpdateUserTagLogic) UpdateUserTag(req *types.UpdateUserTagReq) (*types.UpdateUserTagResp, error) {

    _, err := l.svcCtx.UsersRpc.UpdateUserTag(l.ctx, &users.UpdateUserTagReq{
        Id:          req.Id,
        Name:        req.Name,
        DisplayName: req.DisplayName,
    })
    if err != nil {
        return nil, err
    }

    return &types.UpdateUserTagResp{
        Success: true,
    }, nil
}
