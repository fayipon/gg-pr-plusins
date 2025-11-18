package logic

import (
	"context"
	"users_rpc/internal/model"
	"users_rpc/internal/svc"
)

type GetUserRequest struct {
	Id int64 `json:"id"`
}

type GetUserResponse struct {
	Id       int64  `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *GetUserRequest) (*GetUserResponse, error) {

	u, err := l.svcCtx.UsersModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &GetUserResponse{
		Id:       u.Id,
		Account:  u.Account,
		Password: u.Password,
	}, nil
}
