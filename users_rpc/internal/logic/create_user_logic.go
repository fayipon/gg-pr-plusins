package logic

import (
	"context"
	"users_rpc/internal/model"
	"users_rpc/internal/svc"
)

type CreateUserRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Id int64 `json:"id"`
}

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *CreateUserRequest) (*CreateUserResponse, error) {

	u := &model.Users{
		Account:  req.Account,
		Password: req.Password,
	}

	id, err := l.svcCtx.UsersModel.Insert(l.ctx, u)
	if err != nil {
		return nil, err
	}

	return &CreateUserResponse{Id: id}, nil
}
