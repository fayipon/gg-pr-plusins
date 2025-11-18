package server

import (
	"users_rpc/internal/logic"
	"users_rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

type UsersServer struct {
	ctx *svc.ServiceContext
}

func RegisterUsersServer(group *service.ServiceGroup, ctx *svc.ServiceContext) {
	s := &UsersServer{ctx: ctx}

	group.Add("CreateUser", s.CreateUser)
	group.Add("GetUser", s.GetUser)
	group.Add("UpdateUser", s.UpdateUser)
	group.Add("DeleteUser", s.DeleteUser)
	group.Add("ListUsers", s.ListUsers)
}

func (s *UsersServer) CreateUser(req *logic.CreateUserRequest) (*logic.CreateUserResponse, error) {
	return logic.NewCreateUserLogic(s.ctx).CreateUser(req)
}

func (s *UsersServer) GetUser(req *logic.GetUserRequest) (*logic.GetUserResponse, error) {
	return logic.NewGetUserLogic(s.ctx).GetUser(req)
}

func (s *UsersServer) UpdateUser(req *logic.UpdateUserRequest) (*logic.UpdateUserResponse, error) {
	return logic.NewUpdateUserLogic(s.ctx).UpdateUser(req)
}

func (s *UsersServer) DeleteUser(req *logic.DeleteUserRequest) (*logic.DeleteUserResponse, error) {
	return logic.NewDeleteUserLogic(s.ctx).DeleteUser(req)
}

func (s *UsersServer) ListUsers(req *logic.ListUsersRequest) (*logic.ListUsersResponse, error) {
	return logic.NewListUsersLogic(s.ctx).ListUsers(req)
}
