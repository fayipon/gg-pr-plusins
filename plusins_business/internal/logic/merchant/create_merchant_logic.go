package merchant

import (
    "context"

    "github.com/fayipon/gg-pr-plusins/merchant-service/internal/svc"
    "github.com/fayipon/gg-pr-plusins/merchant-service/internal/types"
)

type CreateMerchantLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewCreateMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMerchantLogic {
    return &CreateMerchantLogic{ctx: ctx, svcCtx: svcCtx}
}

func (l *CreateMerchantLogic) CreateMerchant(req *types.CreateMerchantReq) error {
    _, err := l.svcCtx.DB.Merchant.
        Create().
        SetMerchantCode(req.MerchantCode).
        SetName(req.Name).
        SetContactEmail(req.ContactEmail).
        SetContactPhone(req.ContactPhone).
        SetDomain(req.Domain).
        Save(l.ctx)
    return err
}
