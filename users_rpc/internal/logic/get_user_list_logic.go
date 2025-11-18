package logic

import (
	"context"

	"github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
	"github.com/fayipon/gg-pr-plusins/users_rpc/users"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *users.GetUserListReq) (*users.GetUserListResp, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	size := req.PageSize
	if size <= 0 {
		size = 20
	}
	offset := (page - 1) * size

	list, err := l.svcCtx.UsersModel.List(offset, size, req.Keyword)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.UsersModel.Count(req.Keyword)
	if err != nil {
		return nil, err
	}

	resp := &users.GetUserListResp{
		Total: total,
	}

	for _, u := range list {
		resp.List = append(resp.List, &users.UserItem{
			Id:      uint64(u.Id),
			Account: u.Account,
		})
	}

	return resp, nil
}
