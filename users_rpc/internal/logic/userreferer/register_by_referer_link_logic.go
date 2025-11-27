package userreferer

import (
    "context"
    "time"
	"fmt"
	"strconv"

    "github.com/zeromicro/go-zero/core/logx"

    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/model"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
)

type RegisterByRefererLinkLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewRegisterByRefererLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterByRefererLinkLogic {
    return &RegisterByRefererLinkLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *RegisterByRefererLinkLogic) RegisterByRefererLink(req *users.RegisterByRefererLinkReq) (*users.RegisterByRefererLinkResp, error) {
	
	parentIdStr := req.RefererCode

	parentId, err := strconv.ParseUint(parentIdStr, 10, 64)
	if err != nil {
		return &users.RegisterByRefererLinkResp{
			Success: false,
		}, fmt.Errorf("invalid referer code: %s", parentIdStr)
	}

	parent, err := l.svcCtx.UserRefererModel.FindOneByUserId(l.ctx, parentId)
	if err != nil {
		return &users.RegisterByRefererLinkResp{
			Success: false,
		}, err
	}

    update := &model.UserReferer{
        Id:                parent.Id,
        RegisterCount:     parent.RegisterCount + 1,
        UpdatedAt:         time.Now().Unix(),
    }

    err = l.svcCtx.UserRefererModel.Update(l.ctx, update)
    if err != nil {
        return nil, err
    }

    return &users.RegisterByRefererLinkResp{
        Success: true,
    }, nil
}
