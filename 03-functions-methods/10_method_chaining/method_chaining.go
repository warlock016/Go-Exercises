package method_chaining

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
	return nil
}

// Select sets the columns to select
func (q *QueryBuilder) Select(columns ...string) *QueryBuilder {
	// TODO(human): Implement
	return nil
}

// From sets the table name
func (q *QueryBuilder) From(table string) *QueryBuilder {
	// TODO(human): Implement
	return nil
}

// Where sets the WHERE condition
func (q *QueryBuilder) Where(condition string) *QueryBuilder {
	// TODO(human): Implement
	return nil
}

// OrderBy sets the ORDER BY column
func (q *QueryBuilder) OrderBy(column string) *QueryBuilder {
	// TODO(human): Implement
	return nil
}

// Limit sets the LIMIT value
func (q *QueryBuilder) Limit(n int) *QueryBuilder {
	// TODO(human): Implement
	return nil
}

// Build constructs the final SQL query string
func (q *QueryBuilder) Build() string {
	// TODO(human): Implement
	return ""
}
