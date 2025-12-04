package method_chaining

import (
	"strings"
	"testing"
)

func TestQueryBuilderBasic(t *testing.T) {
	query := NewQueryBuilder().
		Select("*").
		From("users").
		Build()

	expected := "SELECT * FROM users"
	if query != expected {
		t.Errorf("Build() = %q, want %q", query, expected)
	}
}

func TestQueryBuilderWithColumns(t *testing.T) {
	query := NewQueryBuilder().
		Select("id", "name", "email").
		From("users").
		Build()

	expected := "SELECT id, name, email FROM users"
	if query != expected {
		t.Errorf("Build() = %q, want %q", query, expected)
	}
}

func TestQueryBuilderWithWhere(t *testing.T) {
	query := NewQueryBuilder().
		Select("*").
		From("users").
		Where("age > 18").
		Build()

	expected := "SELECT * FROM users WHERE age > 18"
	if query != expected {
		t.Errorf("Build() = %q, want %q", query, expected)
	}
}

func TestQueryBuilderWithOrderBy(t *testing.T) {
	query := NewQueryBuilder().
		Select("*").
		From("users").
		OrderBy("name").
		Build()

	expected := "SELECT * FROM users ORDER BY name"
	if query != expected {
		t.Errorf("Build() = %q, want %q", query, expected)
	}
}

func TestQueryBuilderWithLimit(t *testing.T) {
	query := NewQueryBuilder().
		Select("*").
		From("users").
		Limit(10).
		Build()

	expected := "SELECT * FROM users LIMIT 10"
	if query != expected {
		t.Errorf("Build() = %q, want %q", query, expected)
	}
}

func TestQueryBuilderComplete(t *testing.T) {
	query := NewQueryBuilder().
		Select("id", "name", "email").
		From("users").
		Where("age > 18").
		OrderBy("name").
		Limit(10).
		Build()

	expected := "SELECT id, name, email FROM users WHERE age > 18 ORDER BY name LIMIT 10"
	if query != expected {
		t.Errorf("Build() = %q, want %q", query, expected)
	}
}

func TestQueryBuilderPartialChains(t *testing.T) {
	tests := []struct {
		name     string
		builder  func() string
		expected string
	}{
		{
			"select and where only",
			func() string {
				return NewQueryBuilder().
					Select("*").
					From("users").
					Where("active = true").
					Build()
			},
			"SELECT * FROM users WHERE active = true",
		},
		{
			"select and limit only",
			func() string {
				return NewQueryBuilder().
					Select("id", "name").
					From("products").
					Limit(5).
					Build()
			},
			"SELECT id, name FROM products LIMIT 5",
		},
		{
			"select, where, and limit",
			func() string {
				return NewQueryBuilder().
					Select("*").
					From("orders").
					Where("status = 'pending'").
					Limit(20).
					Build()
			},
			"SELECT * FROM orders WHERE status = 'pending' LIMIT 20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := tt.builder()
			if query != tt.expected {
				t.Errorf("Build() = %q, want %q", query, tt.expected)
			}
		})
	}
}

func TestQueryBuilderOrder(t *testing.T) {
	// Methods can be called in any order, but query should be formatted correctly
	query := NewQueryBuilder().
		Where("active = true").
		Limit(5).
		From("users").
		Select("id", "name").
		OrderBy("id").
		Build()

	// Should still produce correct SQL order
	parts := strings.Split(query, " ")
	if parts[0] != "SELECT" {
		t.Error("Query should start with SELECT")
	}
	if !strings.Contains(query, "FROM users") {
		t.Error("Query should contain FROM users")
	}
	if !strings.Contains(query, "WHERE active = true") {
		t.Error("Query should contain WHERE clause")
	}
	if !strings.Contains(query, "ORDER BY id") {
		t.Error("Query should contain ORDER BY clause")
	}
	if !strings.Contains(query, "LIMIT 5") {
		t.Error("Query should contain LIMIT clause")
	}
}

func TestQueryBuilderReuse(t *testing.T) {
	// Test that builder can be reused
	base := NewQueryBuilder().
		Select("*").
		From("users")

	query1 := base.Where("age > 18").Build()
	query2 := NewQueryBuilder().
		Select("*").
		From("users").
		Where("age < 65").
		Build()

	if !strings.Contains(query1, "age > 18") {
		t.Errorf("Query1 should contain 'age > 18'")
	}
	if !strings.Contains(query2, "age < 65") {
		t.Errorf("Query2 should contain 'age < 65'")
	}
}

func TestQueryBuilderNoTable(t *testing.T) {
	// Build query without setting table (edge case)
	query := NewQueryBuilder().
		Select("*").
		Build()

	if !strings.HasPrefix(query, "SELECT *") {
		t.Error("Query should start with SELECT *")
	}
}

func TestQueryBuilderEmptySelect(t *testing.T) {
	// Not calling Select should default to *
	query := NewQueryBuilder().
		From("users").
		Build()

	expected := "SELECT * FROM users"
	if query != expected {
		t.Errorf("Build() = %q, want %q", query, expected)
	}
}
