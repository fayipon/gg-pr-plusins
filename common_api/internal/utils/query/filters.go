package query

import (
    sq "github.com/Masterminds/squirrel"
    "strings"
)

func ApplyFilters(sb sq.SelectBuilder, filters []FilterItem, allowed map[string]bool) sq.SelectBuilder {
    for _, f := range filters {

        // 无白名单字段一律忽略（安全）
        if !allowed[f.Field] {
            continue
        }

        op := strings.ToLower(f.Op)

        switch op {
        case "=":
            sb = sb.Where(sq.Eq{f.Field: f.Value})

        case "!=":
            sb = sb.Where(sq.NotEq{f.Field: f.Value})

        case ">":
            sb = sb.Where(sq.Gt{f.Field: f.Value})

        case "<":
            sb = sb.Where(sq.Lt{f.Field: f.Value})

        case ">=":
            sb = sb.Where(sq.GtOrEq{f.Field: f.Value})

        case "<=":
            sb = sb.Where(sq.LtOrEq{f.Field: f.Value})

        case "like":
            sb = sb.Where(sq.Like{f.Field: "%" + f.Value.(string) + "%"})

        case "in":
            sb = sb.Where(sq.Eq{f.Field: f.Value})

        case "between":
            vals := f.Value.([]interface{})
            sb = sb.Where(sq.And{
                sq.GtOrEq{f.Field: vals[0]},
                sq.LtOrEq{f.Field: vals[1]},
            })
        }
    }

    return sb
}
