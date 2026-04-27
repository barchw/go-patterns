/*
Package builder implements the Builder pattern.

What is it?
Builder allows constructing complex objects step by step using a fluent API.
Instead of a constructor with many parameters, each method sets one aspect
of the object and returns a reference to the builder.

When to use?
  - When an object has many optional parameters (the "telescoping constructor problem").
  - When the creation process is multi-step and the order of steps may vary.
  - When you want to validate the object's completeness only at the finalization step (Build()).
  - When the same building process should produce different representations.

When NOT to use?
  - When the object has few fields and a simple constructor is readable enough.
  - When all parameters are required and there are no optional combinations.
  - When the cost of maintaining an extra type (Builder) outweighs the readability benefit.

Tips and pitfalls:
  - Place validations in Build(), not in intermediate methods -- this enables call chaining.
  - The builder mutates its state -- do not reuse one instance to build multiple variants.
  - The hasLimit (bool) field lets you distinguish "no limit" from "limit equal to 0".
  - An alternative in Go: Functional Options (WithTimeout(...)) -- better for struct configuration.
*/
package builder

import (
	"errors"
	"fmt"
	"strings"
)

// QueryBuilder builds SQL queries step by step using a fluent API.
type QueryBuilder struct {
	fields    []string
	table     string
	where     []string
	orderBy   string
	limit     int
	hasLimit  bool
}

// NewQueryBuilder creates a new, empty SQL query builder instance.
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

// Select sets the list of fields to retrieve in the query (SELECT clause).
func (q *QueryBuilder) Select(fields ...string) *QueryBuilder {
	q.fields = fields
	return q
}

// From sets the source table name (FROM clause).
func (q *QueryBuilder) From(table string) *QueryBuilder {
	q.table = table
	return q
}

// Where adds a filtering condition (WHERE clause). Multiple calls are joined with the AND operator.
func (q *QueryBuilder) Where(condition string) *QueryBuilder {
	q.where = append(q.where, condition)
	return q
}

// OrderBy sets the field for sorting results (ORDER BY clause).
func (q *QueryBuilder) OrderBy(field string) *QueryBuilder {
	q.orderBy = field
	return q
}

// Limit sets the maximum number of returned rows (LIMIT clause).
func (q *QueryBuilder) Limit(n int) *QueryBuilder {
	q.limit = n
	q.hasLimit = true
	return q
}

// Build finalizes the SQL query construction, validates the required fields, and returns the finished string.
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
