package user

import (
	"context"
	"errors"
	"time"
	"plusins_user_api/internal/svc"
	"plusins_user_api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)


type LoginLogic struct {
	logx.Logger
	ctx context.Context
	svcCtx *svc.ServiceContext
}


func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx: ctx,
		svcCtx: svcCtx,
	}
}


func (l *LoginLogic) Login(req *types.LoginRequest) (*types.LoginResponse, error) {
	if req.Username != "admin" || req.Password != "123456" {
		return nil, errors.New("帐号或密码错误")
	}

	now := time.Now().Unix()
	accessToken, err := l.svcCtx.JwtAuth.GenerateToken(map[string]any{
		"userId": 1,
		"username": req.Username,
		"exp": now + l.svcCtx.Config.Auth.AccessExpire,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		AccessToken: accessToken,
		Expire: now + l.svcCtx.Config.Auth.AccessExpire,
	}, nil
}