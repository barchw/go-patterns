package builder

import "testing"

func TestQueryBuilder_Build(t *testing.T) {
	tests := []struct {
		name     string
		build    func() (string, error)
		expected string
		wantErr  bool
	}{
		{
			name: "simple select",
			build: func() (string, error) {
				return NewQueryBuilder().
					Select("id", "name").
					From("users").
					Build()
			},
			expected: "SELECT id, name FROM users",
		},
		{
			name: "select with where",
			build: func() (string, error) {
				return NewQueryBuilder().
					Select("*").
					From("orders").
					Where("status = 'active'").
					Build()
			},
			expected: "SELECT * FROM orders WHERE status = 'active'",
		},
		{
			name: "select with multiple where clauses",
			build: func() (string, error) {
				return NewQueryBuilder().
					Select("id", "email").
					From("users").
					Where("age > 18").
					Where("active = true").
					Build()
			},
			expected: "SELECT id, email FROM users WHERE age > 18 AND active = true",
		},
		{
			name: "full query",
			build: func() (string, error) {
				return NewQueryBuilder().
					Select("id", "name", "email").
					From("users").
					Where("active = true").
					OrderBy("name").
					Limit(10).
					Build()
			},
			expected: "SELECT id, name, email FROM users WHERE active = true ORDER BY name LIMIT 10",
		},
		{
			name: "select with order and limit only",
			build: func() (string, error) {
				return NewQueryBuilder().
					Select("*").
					From("products").
					OrderBy("price DESC").
					Limit(5).
					Build()
			},
			expected: "SELECT * FROM products ORDER BY price DESC LIMIT 5",
		},
		{
			name: "missing fields",
			build: func() (string, error) {
				return NewQueryBuilder().
					From("users").
					Build()
			},
			wantErr: true,
		},
		{
			name: "missing table",
			build: func() (string, error) {
				return NewQueryBuilder().
					Select("id").
					Build()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.build()
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}
