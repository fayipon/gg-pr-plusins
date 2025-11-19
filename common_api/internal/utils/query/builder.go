package query

import (
    sq "github.com/Masterminds/squirrel"
)

func BuildListQuery(q *QueryEngine) (sq.SelectBuilder, sq.SelectBuilder) {

    sb := sq.Select(q.selectCols...).From(q.table)

    // ---------------------------------------------------------
    // Keyword 搜索
    // ---------------------------------------------------------
    if q.keyword != "" && len(q.keywordCols) > 0 {
        ors := sq.Or{}
        for _, col := range q.keywordCols {
            ors = append(ors, sq.Like{col: "%" + q.keyword + "%"})
        }
        sb = sb.Where(ors)
    }

    // ---------------------------------------------------------
    // 通用 Filter（来自 filters.go）
    // ---------------------------------------------------------
    sb = ApplyFilters(sb, q.filters, q.allowedCols)

    // ---------------------------------------------------------
    // 排序
    // ---------------------------------------------------------
    if q.sortField != "" {
        order := "ASC"
        if q.sortOrder == "DESC" {
            order = "DESC"
        }
        sb = sb.OrderBy(q.sortField + " " + order)
    }

    // ---------------------------------------------------------
    // 分页
    // ---------------------------------------------------------
    offset := (q.page - 1) * q.pageSize
    sb = sb.Offset(uint64(offset)).Limit(uint64(q.pageSize))

    // ---------------------------------------------------------
    // total 查询
    // ---------------------------------------------------------
    cb := sq.Select("COUNT(*)").From(q.table)
    cb = ApplyFilters(cb, q.filters, q.allowedCols)

    return sb, cb
}
