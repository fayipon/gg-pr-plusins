package logic

import (
	"context"
	"users_rpc/internal/svc"
)

type DeleteUserRequest struct {
	Id int64 `json:"id"`
}

type DeleteUserResponse struct {
	Deleted bool `json:"deleted"`
}

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *DeleteUserRequest) (*DeleteUserResponse, error) {

	err := l.svcCtx.UsersModel.Delete(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &DeleteUserResponse{Deleted: true}, nil
}
