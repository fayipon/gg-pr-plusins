package query

import (
    sq "github.com/Masterminds/squirrel"
    "strings"
)

type FilterItem struct {
    Field string      `json:"field"`
    Op    string      `json:"op"`
    Value interface{} `json:"value"`
}

type QueryEngine struct {
    table       string
    selectCols  []string
    filters     []FilterItem
    keyword     string
    sortField   string
    sortOrder   string
    page        int
    pageSize    int
    allowedCols map[string]bool
    keywordCols []string
}

func New(table string, allowed map[string]bool, selectCols []string, keywordCols []string) *QueryEngine {
    return &QueryEngine{
        table:       table,
        allowedCols: allowed,
        selectCols:  selectCols,
        keywordCols: keywordCols,
        page:        1,
        pageSize:    20,
    }
}

// 复制引擎（如果作为 middleware 注入用）
func (q *QueryEngine) Clone() *QueryEngine {
    return &QueryEngine{
        table:       q.table,
        selectCols:  q.selectCols,
        allowedCols: q.allowedCols,
        keywordCols: q.keywordCols,
    }
}

func (q *QueryEngine) Keyword(k string) *QueryEngine {
    q.keyword = k
    return q
}

func (q *QueryEngine) Filters(filters []FilterItem) *QueryEngine {
    q.filters = filters
    return q
}

func (q *QueryEngine) Sort(field, order string) *QueryEngine {
    if q.allowedCols[field] {
        q.sortField = field
        q.sortOrder = strings.ToUpper(order)
    }
    return q
}

func (q *QueryEngine) Page(page, size int) *QueryEngine {
    q.page = page
    q.pageSize = size
    return q
}

// Build() 只负责生成 Squirrel Builders
// 真正的 SQL 生成由 builder.go 完成
func (q *QueryEngine) Build() (sq.SelectBuilder, sq.SelectBuilder) {
    return BuildListQuery(q)
}
