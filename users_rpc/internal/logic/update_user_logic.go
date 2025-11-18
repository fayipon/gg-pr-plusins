package logic

import (
	"context"
	"users_rpc/internal/model"
	"users_rpc/internal/svc"
)

type UpdateUserRequest struct {
	Id       int64  `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

type UpdateUserResponse struct {
	Updated bool `json:"updated"`
}

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *UpdateUserRequest) (*UpdateUserResponse, error) {

	u := &model.Users{
		Id:       req.Id,
		Account:  req.Account,
		Password: req.Password,
	}

	err := l.svcCtx.UsersModel.Update(l.ctx, u)
	if err != nil {
		return nil, err
	}

	return &UpdateUserResponse{Updated: true}, nil
}
