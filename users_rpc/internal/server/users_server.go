package server

import (
    "context"

    logicUser "github.com/fayipon/gg-pr-plusins/users_rpc/internal/logic/user"
    logicUserLevel "github.com/fayipon/gg-pr-plusins/users_rpc/internal/logic/userlevel"
    logicUserGroup "github.com/fayipon/gg-pr-plusins/users_rpc/internal/logic/usergroup"
    logicUserTag "github.com/fayipon/gg-pr-plusins/users_rpc/internal/logic/usertag"
    logicUserReferer "github.com/fayipon/gg-pr-plusins/users_rpc/internal/logic/userreferer"
    logicCommon "github.com/fayipon/gg-pr-plusins/users_rpc/internal/logic/common"
    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type UsersServer struct {
    svcCtx *svc.ServiceContext
    users.UnimplementedUsersServer
}

func NewUsersServer(svcCtx *svc.ServiceContext) *UsersServer {
    return &UsersServer{
        svcCtx: svcCtx,
    }
}

//
// ----------------------
// User
// ----------------------
func (s *UsersServer) GetUserList(ctx context.Context, req *users.GetUserListReq) (*users.GetUserListResp, error) {
    l := logicUser.NewGetUserListLogic(ctx, s.svcCtx)
    return l.GetUserList(req)
}

func (s *UsersServer) GetUser(ctx context.Context, req *users.GetUserReq) (*users.GetUserResp, error) {
    l := logicUser.NewGetUserLogic(ctx, s.svcCtx)
    return l.GetUser(req)
}

func (s *UsersServer) CreateUser(ctx context.Context, req *users.CreateUserReq) (*users.CreateUserResp, error) {
    l := logicUser.NewCreateUserLogic(ctx, s.svcCtx)
    return l.CreateUser(req)
}

//
// ----------------------
// UserLevel
// ----------------------
func (s *UsersServer) CreateUserLevel(ctx context.Context, req *users.CreateUserLevelReq) (*users.CreateUserLevelResp, error) {
    l := logicUserLevel.NewCreateUserLevelLogic(ctx, s.svcCtx)
    return l.CreateUserLevel(req)
}

func (s *UsersServer) GetUserLevel(ctx context.Context, req *users.GetUserLevelReq) (*users.GetUserLevelResp, error) {
    l := logicUserLevel.NewGetUserLevelLogic(ctx, s.svcCtx)
    return l.GetUserLevel(req)
}

func (s *UsersServer) UpdateUserLevel(ctx context.Context, req *users.UpdateUserLevelReq) (*users.UpdateUserLevelResp, error) {
    l := logicUserLevel.NewUpdateUserLevelLogic(ctx, s.svcCtx)
    return l.UpdateUserLevel(req)
}

func (s *UsersServer) DeleteUserLevel(ctx context.Context, req *users.DeleteUserLevelReq) (*users.DeleteUserLevelResp, error) {
    l := logicUserLevel.NewDeleteUserLevelLogic(ctx, s.svcCtx)
    return l.DeleteUserLevel(req)
}

func (s *UsersServer) GetUserLevelList(ctx context.Context, req *users.GetUserLevelListReq) (*users.GetUserLevelListResp, error) {
    l := logicUserLevel.NewGetUserLevelListLogic(ctx, s.svcCtx)
    return l.GetUserLevelList(req)
}

//
// ----------------------
// UserGroup
// ----------------------
func (s *UsersServer) CreateUserGroup(ctx context.Context, req *users.CreateUserGroupReq) (*users.CreateUserGroupResp, error) {
    l := logicUserGroup.NewCreateUserGroupLogic(ctx, s.svcCtx)
    return l.CreateUserGroup(req)
}

func (s *UsersServer) GetUserGroup(ctx context.Context, req *users.GetUserGroupReq) (*users.GetUserGroupResp, error) {
    l := logicUserGroup.NewGetUserGroupLogic(ctx, s.svcCtx)
    return l.GetUserGroup(req)
}

func (s *UsersServer) UpdateUserGroup(ctx context.Context, req *users.UpdateUserGroupReq) (*users.UpdateUserGroupResp, error) {
    l := logicUserGroup.NewUpdateUserGroupLogic(ctx, s.svcCtx)
    return l.UpdateUserGroup(req)
}

func (s *UsersServer) DeleteUserGroup(ctx context.Context, req *users.DeleteUserGroupReq) (*users.DeleteUserGroupResp, error) {
    l := logicUserGroup.NewDeleteUserGroupLogic(ctx, s.svcCtx)
    return l.DeleteUserGroup(req)
}

func (s *UsersServer) GetUserGroupList(ctx context.Context, req *users.GetUserGroupListReq) (*users.GetUserGroupListResp, error) {
    l := logicUserGroup.NewGetUserGroupListLogic(ctx, s.svcCtx)
    return l.GetUserGroupList(req)
}

//
// ----------------------
// UserTags
// ----------------------

func (s *UsersServer) CreateUserTag(ctx context.Context, req *users.CreateUserTagReq) (*users.CreateUserTagResp, error) {
    l := logicUserTag.NewCreateUserTagLogic(ctx, s.svcCtx)
    return l.CreateUserTag(req)
}

func (s *UsersServer) GetUserTag(ctx context.Context, req *users.GetUserTagReq) (*users.GetUserTagResp, error) {
    l := logicUserTag.NewGetUserTagLogic(ctx, s.svcCtx)
    return l.GetUserTag(req)
}

func (s *UsersServer) UpdateUserTag(ctx context.Context, req *users.UpdateUserTagReq) (*users.UpdateUserTagResp, error) {
    l := logicUserTag.NewUpdateUserTagLogic(ctx, s.svcCtx)
    return l.UpdateUserTag(req)
}

func (s *UsersServer) DeleteUserTag(ctx context.Context, req *users.DeleteUserTagReq) (*users.DeleteUserTagResp, error) {
    l := logicUserTag.NewDeleteUserTagLogic(ctx, s.svcCtx)
    return l.DeleteUserTag(req)
}

func (s *UsersServer) GetUserTagList(ctx context.Context, req *users.GetUserTagListReq) (*users.GetUserTagListResp, error) {
    l := logicUserTag.NewGetUserTagListLogic(ctx, s.svcCtx)
    return l.GetUserTagList(req)
}

//
// ----------------------
// UserReferer
// ----------------------

func (s *UsersServer) CreateUserReferer(ctx context.Context, req *users.CreateUserRefererReq) (*users.CreateUserRefererResp, error) {
    l := logicUserReferer.NewCreateUserRefererLogic(ctx, s.svcCtx)
    return l.CreateUserReferer(req)
}

func (s *UsersServer) GetUserReferer(ctx context.Context, req *users.GetUserRefererReq) (*users.GetUserRefererResp, error) {
    l := logicUserReferer.NewGetUserRefererLogic(ctx, s.svcCtx)
    return l.GetUserReferer(req)
}

func (s *UsersServer) UpdateUserReferer(ctx context.Context, req *users.UpdateUserRefererReq) (*users.UpdateUserRefererResp, error) {
    l := logicUserReferer.NewUpdateUserRefererLogic(ctx, s.svcCtx)
    return l.UpdateUserReferer(req)
}

func (s *UsersServer) DeleteUserReferer(ctx context.Context, req *users.DeleteUserRefererReq) (*users.DeleteUserRefererResp, error) {
    l := logicUserReferer.NewDeleteUserRefererLogic(ctx, s.svcCtx)
    return l.DeleteUserReferer(req)
}

func (s *UsersServer) GetUserRefererList(ctx context.Context, req *users.GetUserRefererListReq) (*users.GetUserRefererListResp, error) {
    l := logicUserReferer.NewGetUserRefererListLogic(ctx, s.svcCtx)
    return l.GetUserRefererList(req)
}

func (s *UsersServer) GenerateRefererLink(ctx context.Context, req *users.GenerateRefererLinkReq) (*users.GenerateRefererLinkResp, error) {
    l := logicUserReferer.NewGenerateRefererLinkLogic(ctx, s.svcCtx)
    return l.GenerateRefererLink(req)
}

func (s *UsersServer) RegisterByRefererLink(ctx context.Context, req *users.RegisterByRefererLinkReq) (*users.RegisterByRefererLinkResp, error) {
    l := logicUserReferer.NewRegisterByRefererLinkLogic(ctx, s.svcCtx)
    return l.RegisterByRefererLink(req)
}

//
// ----------------------
// Level Batch
// ----------------------
func (s *UsersServer) GetLevelsBatch(ctx context.Context, in *users.LevelBatchReq) (*users.LevelBatchResp, error) {
    l := logicCommon.NewGetLevelsBatchLogic(ctx, s.svcCtx)
    return l.GetLevelsBatch(in)
}
