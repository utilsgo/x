package context

import (
	"context"
	"testing"
)

type contextKey struct{}

func getValue(ctx context.Context) bool {
	v := ctx.Value(contextKey{})
	_ = v
	return true
}

func BenchmarkWithValue(b *testing.B) {
	parent := context.Background()

	b.Run("WithValue with reflect type comparable check", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ctx := context.WithValue(parent, contextKey{}, nil)
			getValue(ctx)
		}
	})

	b.Run("WithValue without reflect type comparable check", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ctx := WithValue(parent, contextKey{}, nil)
			getValue(ctx)
		}
	})
}
