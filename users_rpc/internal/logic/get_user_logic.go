package logic

import (
    "context"
    "database/sql"

    sq "github.com/Masterminds/squirrel"
    "github.com/fayipon/gg-pr-plusins/users_rpc/internal/svc"
    "github.com/fayipon/gg-pr-plusins/users_rpc/users"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
    return &GetUserLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetUserLogic) GetUser(req *users.GetUserReq) (*users.GetUserResp, error) {

    db := l.svcCtx.DB

    // Build SQL
    qb := sq.Select(
        "id",
        "account",
        "status",
        "level_id",
        "created_at",
    ).From("users").
        Where(sq.Eq{"id": req.Id})

    sqlStr, args, _ := qb.ToSql()

    // 结构体 mapping
    var row struct {
        Id        uint64       `db:"id"`
        Account   string       `db:"account"`
        Status    int64        `db:"status"`
        LevelId   int64        `db:"level_id"`
        CreatedAt sql.NullTime `db:"created_at"`
    }

    // go-zero 推荐方式：QueryRowCtx + Scan
    err := db.QueryRowCtx(l.ctx, func(scan func(dest ...any) error) error {
        return scan(
            &row.Id,
            &row.Account,
            &row.Status,
            &row.LevelId,
            &row.CreatedAt,
        )
    }, sqlStr, args...)

    if err != nil {
        // 没有找到
        if err == sql.ErrNoRows {
            return nil, err
        }
        return nil, err
    }

    return &users.GetUserResp{
        Id:        row.Id,
        Account:   row.Account,
        Status:    row.Status,
        LevelId:   row.LevelId,
        CreatedAt: row.CreatedAt.Time.Unix(),
    }, nil
}
