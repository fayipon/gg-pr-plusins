package userlevel

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "common_api/internal/utils/errorx"
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

func (l *GetUserLevelLogic) GetUserLevel(req *types.GetUserLevelReq) (*types.GetUserLevelResp, error) {

    // Debug
    l.Infof(">>> GetUserLevel Logic received ID = %d", req.Id)

    if req.Id == 0 {
        return nil, errorx.NewCodeError(l.ctx, errorx.ErrInvalidParams, "id must be > 0")
    }

    rpcResp, err := l.svcCtx.UsersRpc.GetUserLevel(
        l.ctx,
        &users.GetUserLevelReq{
            Id: req.Id,
        },
    )

    if err != nil {
        return nil, err
    }

    return &types.GetUserLevelResp{
        Id:          rpcResp.Id,
        Name:        rpcResp.Name,
        DisplayName: rpcResp.DisplayName,
        Setting:     rpcResp.Setting,
        CreatedAt:   rpcResp.CreatedAt,
        UpdatedAt:   rpcResp.UpdatedAt,
    }, nil
}
