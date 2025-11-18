package server

import (
	"context"

	"users_rpc/internal/logic"
	"users_rpc/internal/svc"
	"users_rpc/users"
)

type UsersServer struct {
	svcCtx *svc.ServiceContext
	users.UnimplementedUsersServer
}

func NewUsersServer(ctx *svc.ServiceContext) *UsersServer {
	return &UsersServer{
		svcCtx: ctx,
	}
}

func (s *UsersServer) CreateUser(ctx context.Context, in *users.CreateUserReq) (*users.CreateUserResp, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *UsersServer) GetUser(ctx context.Context, in *users.GetUserReq) (*users.GetUserResp, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

func (s *UsersServer) UpdateUser(ctx context.Context, in *users.UpdateUserReq) (*users.UpdateUserResp, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *UsersServer) DeleteUser(ctx context.Context, in *users.DeleteUserReq) (*users.DeleteUserResp, error) {
	l := logic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}

func (s *UsersServer) ListUser(ctx context.Context, in *users.ListUserReq) (*users.ListUserResp, error) {
	l := logic.NewListUserLogic(ctx, s.svcCtx)
	return l.ListUser(in)
}
