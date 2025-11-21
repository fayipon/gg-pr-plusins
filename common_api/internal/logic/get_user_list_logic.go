package logic

import (
    "context"

    "common_api/internal/svc"
    "common_api/internal/types"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
    return &GetUserListLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (*types.UserListResp, error) {

    rpcReq := &users.GetUserListReq{
        Page:     int32(req.Page),
        PageSize: int32(req.PageSize),
        Keyword:  req.Keyword,
    }

    rpcResp, err := l.svcCtx.UsersRpc.GetUserList(l.ctx, rpcReq)
    if err != nil {
        logx.Errorf("UsersRpc.GetUserList error: %+v", err)
        return nil, err
    }

    // 🔥 修正成 []*types.UserListItem
    list := make([]*types.UserListItem, 0, len(rpcResp.List))

    // 用来收集所有 level_id
    levelIds := make([]uint64, 0, len(rpcResp.List))

    // 先复制 user list
    for _, item := range rpcResp.List {

        list = append(list, &types.UserListItem{
            Id:        item.Id,
            Account:   item.Account,
            Status:    item.Status,
            LevelId:   item.LevelId,
            CreatedAt: item.CreatedAt,
        })

        if item.LevelId > 0 {
            levelIds = append(levelIds, item.LevelId)
        }
    }

    //
    // 去重 level_id
    //
    uniq := make(map[uint64]struct{}, len(levelIds))
    finalLevelIDs := make([]uint64, 0, len(levelIds))
    for _, id := range levelIds {
        if _, exists := uniq[id]; !exists {
            uniq[id] = struct{}{}
            finalLevelIDs = append(finalLevelIDs, id)
        }
    }

    //
    // 批次从 users_rpc 取等级资料
    //
    levelMap := make(map[uint64]*users.UserLevelInfo)

    if len(finalLevelIDs) > 0 {
        lvResp, err := l.svcCtx.UsersRpc.GetLevelsBatch(
            l.ctx,
            &users.LevelBatchReq{Ids: finalLevelIDs},
        )
        if err != nil {
            logx.Errorf("UsersRpc.GetLevelsBatch error: %+v", err)
        } else {
            for _, lv := range lvResp.List {
                levelMap[lv.Id] = lv
            }
        }
    }

    //
    // 把等级补进 UserListItem
    //
    for i := range list {
        if info, ok := levelMap[list[i].LevelId]; ok && info != nil {
            list[i].UserLevel = &types.UserLevelInfo{
                Id:          info.Id,
                Name:        info.Name,
                DisplayName: info.DisplayName,
            }
        }
    }

    // 返回
    return &types.UserListResp{
        Total: rpcResp.Total,
        List:  list, // 🔥 类型完全正确：[]*UserListItem
    }, nil
}
