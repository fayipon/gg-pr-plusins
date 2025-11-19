package logic

import (
	"context"

	"common_api/internal/svc"
	"common_api/internal/types"
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

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (*types.UserListResp, error) {

	// ------------------------------------------------------
	// Step 1: 准备 RPC 查询参数
	// ------------------------------------------------------
	rpcReq := &users.GetUserListReq{
		Page:      int32(req.Page),
		PageSize:  int32(req.PageSize),
		Keyword:   req.Keyword,
		SortField: req.SortField,
		SortOrder: req.SortOrder,
	}

	// 转换 Filter 到 RPC 所需格式
	if len(req.Filters) > 0 {
		rpcReq.Filters = make([]*users.FilterItem, 0, len(req.Filters))
		for _, f := range req.Filters {
			rpcReq.Filters = append(rpcReq.Filters, &users.FilterItem{
				Field: f.Field,
				Op:    f.Op,
				Value: toRpcValue(f.Value),
			})
		}
	}

	// ------------------------------------------------------
	// Step 2: 调用 RPC
	// ------------------------------------------------------
	rpcResp, err := l.svcCtx.UsersRpc.GetUserList(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	// ------------------------------------------------------
	// Step 3: 组装 API Response
	// ------------------------------------------------------
	resp := &types.UserListResp{
		Total: rpcResp.Total,
		List:  make([]types.UserListItem, 0, len(rpcResp.List)),
	}

	for _, item := range rpcResp.List {
		resp.List = append(resp.List, types.UserListItem{
			Id:      item.Id,
			Account: item.Account,
			Status:    item.Status,
			LevelId:   item.LevelId,
			CreatedAt: item.CreatedAt,
		})
	}

	return resp, nil
}

// Value 转换助手（interface{} → RPC 可接受类型）
func toRpcValue(v interface{}) *users.Value {
	switch val := v.(type) {

	case string:
		return &users.Value{Value: &users.Value_StrVal{StrVal: val}}

	case float64:
		return &users.Value{Value: &users.Value_NumVal{NumVal: val}}

	case int:
		return &users.Value{Value: &users.Value_NumVal{NumVal: float64(val)}}

	case []interface{}:
		arr := make([]string, 0)
		for _, x := range val {
			arr = append(arr, x.(string))
		}
		return &users.Value{Value: &users.Value_StrArray{StrArray: &users.StringArray{Values: arr}}}

	default:
		return &users.Value{Value: &users.Value_StrVal{StrVal: ""}}
	}
}
