package userreferer

import (
    "context"
    "fmt"

    "github.com/zeromicro/go-zero/core/logx"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type GenerateRefererLinkLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGenerateRefererLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateRefererLinkLogic {
    return &GenerateRefererLinkLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GenerateRefererLinkLogic) GenerateRefererLink(req *users.GenerateRefererLinkReq) (*users.GenerateRefererLinkResp, error) {

    link := fmt.Sprintf("%s/register?ref=%d", l.svcCtx.Config.AppUrl, req.UserId)

    return &users.GenerateRefererLinkResp{
        RefererLink: link,
    }, nil
}
