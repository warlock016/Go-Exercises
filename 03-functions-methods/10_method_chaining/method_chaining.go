package method_chaining

import (
	"strconv"
	"strings"
)

// QueryBuilder builds SQL SELECT queries using method chaining
type QueryBuilder struct {
	table   string
	columns []string
	where   string
	orderBy string
	limit   int
}

// NewQueryBuilder creates a new QueryBuilder
func NewQueryBuilder() *QueryBuilder {
	// TODO(human): Implement
	return &QueryBuilder{
		table:   "",
		columns: make([]string, 0),
		where:   "",
		orderBy: "",
		limit:   0,
	}
}

// Select sets the columns to select
func (q *QueryBuilder) Select(columns ...string) *QueryBuilder {
	// TODO(human): Implement
	cols := make([]string, 0, len(columns))
	if len(q.columns) != 0 {
		cols = append(cols, q.columns...)
	}
	cols = append(cols, columns...)

	q.columns = cols
	return q
}

// From sets the table name
func (q *QueryBuilder) From(table string) *QueryBuilder {
	// TODO(human): Implement
	q.table = table
	return q
}

// Where sets the WHERE condition
func (q *QueryBuilder) Where(condition string) *QueryBuilder {
	// TODO(human): Implement
	q.where = condition
	return q
}

// OrderBy sets the ORDER BY column
func (q *QueryBuilder) OrderBy(column string) *QueryBuilder {
	// TODO(human): Implement
	q.orderBy = column
	return q
}

// Limit sets the LIMIT value
func (q *QueryBuilder) Limit(n int) *QueryBuilder {
	// TODO(human): Implement
	q.limit = n
	return q
}

// Build constructs the final SQL query string
func (q *QueryBuilder) Build() string {
	// TODO(human): Implement
	var result strings.Builder

	result.WriteString("SELECT ")

	if len(q.columns) == 0 {
		result.WriteString("*")
	} else {
		for i, v := range q.columns {
			result.WriteString(v)
			if i != len(q.columns)-1 {
				result.WriteString(", ")
			}
		}
	}

	if q.table != "" {
		result.WriteString(" FROM ")
		result.WriteString(q.table)
	}

	if q.where != "" {
		result.WriteString(" WHERE ")
		result.WriteString(q.where)
	}

	if q.orderBy != "" {
		result.WriteString(" ORDER BY ")
		result.WriteString(q.orderBy)
	}

	if q.limit != 0 {
		result.WriteString(" LIMIT ")
		lim := strconv.FormatInt(int64(q.limit), 10)
		result.WriteString(lim)
	}
	return result.String()
}
