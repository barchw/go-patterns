package builder

import (
	"errors"
	"fmt"
	"strings"
)

type QueryBuilder struct {
	fields    []string
	table     string
	where     []string
	orderBy   string
	limit     int
	hasLimit  bool
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (q *QueryBuilder) Select(fields ...string) *QueryBuilder {
	q.fields = fields
	return q
}

func (q *QueryBuilder) From(table string) *QueryBuilder {
	q.table = table
	return q
}

func (q *QueryBuilder) Where(condition string) *QueryBuilder {
	q.where = append(q.where, condition)
	return q
}

func (q *QueryBuilder) OrderBy(field string) *QueryBuilder {
	q.orderBy = field
	return q
}

func (q *QueryBuilder) Limit(n int) *QueryBuilder {
	q.limit = n
	q.hasLimit = true
	return q
}

func (q *QueryBuilder) Build() (string, error) {
	if len(q.fields) == 0 {
		return "", errors.New("at least one field is required")
	}
	if q.table == "" {
		return "", errors.New("table name is required")
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "SELECT %s FROM %s", strings.Join(q.fields, ", "), q.table)

	for i, w := range q.where {
		if i == 0 {
			fmt.Fprintf(&sb, " WHERE %s", w)
		} else {
			fmt.Fprintf(&sb, " AND %s", w)
		}
	}

	if q.orderBy != "" {
		fmt.Fprintf(&sb, " ORDER BY %s", q.orderBy)
	}

	if q.hasLimit {
		fmt.Fprintf(&sb, " LIMIT %d", q.limit)
	}

	return sb.String(), nil
}
