package logic

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"

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
	db := l.svcCtx.DB

	// pagination
	page := req.Page
	if page <= 0 {
		page = 1
	}
	size := req.PageSize
	if size <= 0 {
		size = 20
	}
	offset := (page - 1) * size

	// base select
	qb := sq.Select(
		"id",
		"account",
		"status",
		"level_id",
		"created_at",
	).From("users")

	// keyword
	if req.Keyword != "" {
		kw := "%" + req.Keyword + "%"
		qb = qb.Where(sq.Like{"account": kw})
	}

	// filters
	for _, f := range req.Filters {
		field := f.Field
		op := f.Op

		switch val := f.Value.Value.(type) {
		case *users.Value_StrVal:
			qb = qb.Where(fmt.Sprintf("%s %s ?", field, op), val.StrVal)

		case *users.Value_NumVal:
			qb = qb.Where(fmt.Sprintf("%s %s ?", field, op), val.NumVal)

		case *users.Value_StrArray:
			qb = qb.Where(sq.Eq{field: val.StrArray.Values})
		}
	}

	// sort
	if req.SortField != "" {
		order := "ASC"
		if req.SortOrder == "DESC" {
			order = "DESC"
		}
		qb = qb.OrderBy(req.SortField + " " + order)
	}

	// pagination
	qb = qb.Offset(uint64(offset)).Limit(uint64(size))

	// build list sql
	listSQL, listArgs, _ := qb.ToSql()
	logx.Infof("LIST SQL = %s args=%v", listSQL, listArgs)

	// struct for scanning
	type userRow struct {
		Id        uint64        `db:"id"`
		Account   string        `db:"account"`
		Status    int64         `db:"status"`
		LevelId     int64         `db:"level_id"`
		CreatedAt sql.NullTime  `db:"created_at"`
	}

	var rows []userRow

	// ✔ 正确 go-zero 写法
	err := db.QueryRowsCtx(l.ctx, &rows, listSQL, listArgs...)
	if err != nil {
		return nil, err
	}

	resp := &users.GetUserListResp{
		List: make([]*users.UserItem, 0, len(rows)),
	}

	for _, r := range rows {
		resp.List = append(resp.List, &users.UserItem{
			Id:        r.Id,
			Account:   r.Account,
			Status:    r.Status,
			LevelId:   r.LevelId,
			CreatedAt: r.CreatedAt.Time.Unix(),
		})
	}

	// ------------------------
	// count sql
	// ------------------------
	cb := sq.Select("COUNT(*)").From("users")

	if req.Keyword != "" {
		kw := "%" + req.Keyword + "%"
		cb = cb.Where(sq.Like{"account": kw})
	}

	countSQL, cArgs, _ := cb.ToSql()

	var total int64

	err = db.QueryRowCtx(l.ctx, &total, countSQL, cArgs...)
	if err != nil {
		return nil, err
	}

	resp.Total = total
	return resp, nil
}
