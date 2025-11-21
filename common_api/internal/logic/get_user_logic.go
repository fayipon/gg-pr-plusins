package logic

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"

    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
    return &GetUserLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetUserLogic) GetUser(in *types.GetUserReq) (*types.GetUserResp, error) {

    rpcResp, err := l.svcCtx.UsersRpc.GetUser(l.ctx, &users.GetUserReq{
        Id: in.Id,
    })
    if err != nil {
        logx.Errorf("[GetUser] RPC error: %+v", err)
        return nil, err

    }

    resp := &types.GetUserResp{
        Id:              rpcResp.Id,
        Account:         rpcResp.Account,
        LevelId:         rpcResp.LevelId,
        GroupId:         rpcResp.GroupId,
        EmailVerifiedAt: rpcResp.EmailVerifiedAt,
        MobileVerifiedAt: rpcResp.MobileVerifiedAt,
        KycVerifiedAt:   rpcResp.KycVerifiedAt,
        ParentId:        rpcResp.ParentId,
        ParentTree:      rpcResp.ParentTree,
        Depth:           rpcResp.Depth,
        RefererId:       rpcResp.RefererId,
        Status:          rpcResp.Status,
        CreatedAt:       rpcResp.CreatedAt,
        UpdatedAt:       rpcResp.UpdatedAt,
    }

    // user level optional
    if rpcResp.UserLevel != nil {
        resp.UserLevel = &types.UserLevelInfo{
            Id:          rpcResp.UserLevel.Id,
            Name:        rpcResp.UserLevel.Name,
            DisplayName: rpcResp.UserLevel.DisplayName,
        }
    }

    return resp, nil
}
