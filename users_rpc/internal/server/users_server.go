package server

import (
	"context"

	"github.com/fayipon/gg-pr-plusins/users_rpc/internal/logic"
	"github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
	"github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type UsersServer struct {
	svcCtx *svc.ServiceContext
	users.UnimplementedUsersServer
}

func NewUsersServer(svcCtx *svc.ServiceContext) *UsersServer {
	return &UsersServer{svcCtx: svcCtx}
}

func (s *UsersServer) GetUserList(ctx context.Context, req *users.GetUserListReq) (*users.GetUserListResp, error) {
	l := logic.NewGetUserListLogic(ctx, s.svcCtx)
	return l.GetUserList(req)
}

func (s *UsersServer) GetUser(ctx context.Context, req *users.GetUserReq) (*users.GetUserResp, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(req)
}

func (s *UsersServer) CreateUser(ctx context.Context, req *users.CreateUserReq) (*users.CreateUserResp, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(req)
}

func (s *UsersServer) GetLevelsBatch(ctx context.Context, in *users.LevelBatchReq) (*users.LevelBatchResp, error) {
    l := logic.NewGetLevelsBatchLogic(ctx, s.svcCtx)
    return l.GetLevelsBatch(in)
}