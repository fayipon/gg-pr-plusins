package sysapi

import (
    "context"

    "github.com/fayipon/gg-pr-plusins/example/internal/svc"
    "github.com/fayipon/gg-pr-plusins/example/internal/types"
)

type VersionLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VersionLogic {
    return &VersionLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *VersionLogic) Version() (*types.VersionResp, error) {
    return &types.VersionResp{
        Version: "0.0.1",
    }, nil
}
