package logic

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"github.com/fayipon/gg-pr-plusins/plusins_business/internal/svc"
	"github.com/fayipon/gg-pr-plusins/plusins_business/desc/types"
)

type RegisterBusinessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterBusinessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterBusinessLogic {
	return &RegisterBusinessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterBusinessLogic) RegisterBusiness(req *types.RegisterBusinessReq) (*types.RegisterBusinessResp, error) {
	token := generateToken()

	return &types.RegisterBusinessResp{
		Id:       1,
		Name:     req.Name,
		Token:    token,
		Region:   req.Region,
		Currency: req.Currency,
		Locale:   req.Locale,
		Status:   "active",
	}, nil
}

func generateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}
