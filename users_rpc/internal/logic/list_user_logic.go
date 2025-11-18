package logic

import (
	"context"
	"users_rpc/internal/svc"
)

type ListUsersRequest struct{}

type ListUsersResponse struct {
	List []UserItem `json:"list"`
}

type UserItem struct {
	Id       int64  `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

type ListUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUsersLogic {
	return &ListUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUsersLogic) ListUsers(req *ListUsersRequest) (*ListUsersResponse, error) {

	users, err := l.svcCtx.UsersModel.List(l.ctx)
	if err != nil {
		return nil, err
	}

	var resp []UserItem
	for _, u := range users {
		resp = append(resp, UserItem{
			Id:       u.Id,
			Account:  u.Account,
			Password: u.Password,
		})
	}

	return &ListUsersResponse{List: resp}, nil
}
