package server

import (
	"context"

	"github.com/fayipon/gg-pr-plusins/users_rpc/internal/logic"
	"github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
	"github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type UsersServer struct {
	users.UnimplementedUsersServer
	ctx *svc.ServiceContext
}

func NewUsersServer(ctx *svc.ServiceContext) *UsersServer {
	return &UsersServer{
		ctx: ctx,
	}
}

func (s *UsersServer) CreateUser(ctx context.Context, req *users.CreateUserReq) (*users.CreateUserResp, error) {
	l := logic.NewCreateUserLogic(ctx, s.ctx)
	return l.CreateUser(req)
}

func (s *UsersServer) GetUser(ctx context.Context, req *users.GetUserReq) (*users.GetUserResp, error) {
	l := logic.NewGetUserLogic(ctx, s.ctx)
	return l.GetUser(req)
}

func (s *UsersServer) GetUserList(ctx context.Context, req *users.GetUserListReq) (*users.GetUserListResp, error) {
	l := logic.NewGetUserListLogic(ctx, s.ctx)
	return l.GetUserList(req)
}