package logic

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/fayipon/gg-pr-plusins/users_rpc/internal/model"
	"github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
	"github.com/fayipon/gg-pr-plusins/users_rpc/users"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *users.CreateUserReq) (*users.CreateUserResp, error) {

	exist, err := l.svcCtx.UsersModel.FindByAccount(req.Account)
	if err != nil {
		return nil, err
	}
	if exist != nil {
		return nil, fmt.Errorf("account already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	data := &model.Users{
		Account:  req.Account,
		Password: string(hash),
	}

	id, err := l.svcCtx.UsersModel.Insert(data)
	if err != nil {
		return nil, err
	}

	return &users.CreateUserResp{
		Id:      uint64(id),
		Account: req.Account,
	}, nil
}
