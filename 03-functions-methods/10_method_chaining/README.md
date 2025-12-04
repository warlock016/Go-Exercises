# Exercise 10: Method Chaining (Builder Pattern)

**Learning Goal:** Master fluent interfaces using method chaining

---

## 📝 Problem Description

Method chaining creates fluent APIs where methods return the receiver, allowing calls to be chained:
`builder.SetA().SetB().SetC().Build()`

This pattern is used for:
- Builders (SQL queries, HTTP requests)
- Configuration objects
- Fluent interfaces
- DSLs (Domain Specific Languages)

---

## 🎯 Type & Method Signatures

```go
type QueryBuilder struct {
    table      string
    columns    []string
    where      string
    orderBy    string
    limit      int
}

func NewQueryBuilder() *QueryBuilder
func (q *QueryBuilder) Select(columns ...string) *QueryBuilder
func (q *QueryBuilder) From(table string) *QueryBuilder
func (q *QueryBuilder) Where(condition string) *QueryBuilder
func (q *QueryBuilder) OrderBy(column string) *QueryBuilder
func (q *QueryBuilder) Limit(n int) *QueryBuilder
func (q *QueryBuilder) Build() string
```

---

## 📖 Examples

```go
query := NewQueryBuilder().
    Select("id", "name", "email").
    From("users").
    Where("age > 18").
    OrderBy("name").
    Limit(10).
    Build()

// "SELECT id, name, email FROM users WHERE age > 18 ORDER BY name LIMIT 10"

// Minimal query
query2 := NewQueryBuilder().
    Select("*").
    From("products").
    Build()
// "SELECT * FROM products"
```

---

## 📋 Instructions

1. Define `QueryBuilder` struct with fields
2. Implement `NewQueryBuilder()` constructor
3. Implement chainable methods that return `*QueryBuilder`
4. Implement `Build()` to generate SQL string
5. Run tests with `go test -v`

---

## 💡 Hints

<details>
<summary>Basic Concept</summary>

Pointer receivers return `*Type` for chaining:
```go
func (b *Builder) Set(v int) *Builder {
    b.value = v
    return b  // Return self for chaining
}
```

</details>

<details>
<summary>Complete Solution</summary>

```go
package method_chaining

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	table   string
	columns []string
	where   string
	orderBy string
	limit   int
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (q *QueryBuilder) Select(columns ...string) *QueryBuilder {
	q.columns = columns
	return q
}

func (q *QueryBuilder) From(table string) *QueryBuilder {
	q.table = table
	return q
}

func (q *QueryBuilder) Where(condition string) *QueryBuilder {
	q.where = condition
	return q
}

func (q *QueryBuilder) OrderBy(column string) *QueryBuilder {
	q.orderBy = column
	return q
}

func (q *QueryBuilder) Limit(n int) *QueryBuilder {
	q.limit = n
	return q
}

func (q *QueryBuilder) Build() string {
	parts := []string{}

	// SELECT
	cols := "*"
	if len(q.columns) > 0 {
		cols = strings.Join(q.columns, ", ")
	}
	parts = append(parts, "SELECT "+cols)

	// FROM
	if q.table != "" {
		parts = append(parts, "FROM "+q.table)
	}

	// WHERE
	if q.where != "" {
		parts = append(parts, "WHERE "+q.where)
	}

	// ORDER BY
	if q.orderBy != "" {
		parts = append(parts, "ORDER BY "+q.orderBy)
	}

	// LIMIT
	if q.limit > 0 {
		parts = append(parts, fmt.Sprintf("LIMIT %d", q.limit))
	}

	return strings.Join(parts, " ")
}
```

</details>

---

## 🤔 Think About

1. Why use pointer receivers for chaining?
2. What are the benefits of fluent APIs?
3. How does this pattern improve readability?
4. What are the downsides of method chaining?

---

## 🎓 What This Teaches

- **Builder pattern**: Step-by-step object construction
- **Method chaining**: Fluent interfaces
- **Pointer receivers**: Returning self for chaining
- **API design**: Creating readable, expressive APIs
- **DSLs**: Domain-specific languages in Go

---

**Tier:** 3 - Integration
**Estimated Time:** 40-50 minutes
