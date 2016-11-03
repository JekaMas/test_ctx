package service

import (
	"testing"
	"context"
	"test_ctx/service/base"
)

type AMock struct {
	base.BaseService
}

func (this *AMock) DoA() int {
	return 0
}

func BenchmarkFromContext(b *testing.B) {
	var key interface{} = "id5"
	ctx := context.Background()
	ctx = context.WithValue(ctx, "id1", 1)
	ctx = context.WithValue(ctx, "id2", 1)
	ctx = context.WithValue(ctx, "id3", 1)
	ctx = context.WithValue(ctx, "id4", 1)
	ctx = context.WithValue(ctx, "id5", 1)
	ctx = context.WithValue(ctx, "id6", 1)
	ctx = context.WithValue(ctx, "id7", 1)
	ctx = context.WithValue(ctx, "id8", 1)
	ctx = context.WithValue(ctx, "id9", 1)
	ctx = context.WithValue(ctx, "id10", 1)


	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx.Value(key)
		}
	})
}

func BenchmarkCopy(b *testing.B) {
	ctx := context.Background()
	val := &AMock{}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			copy(ctx, val)
		}
	})
}

func BenchmarkAFromContext(b *testing.B) {
	factory := &Factory{
		A: &AMock{},
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "id1", 1)
	ctx = context.WithValue(ctx, "id2", 1)
	ctx = context.WithValue(ctx, "id3", 1)
	ctx = context.WithValue(ctx, "id4", 1)
	ctx = context.WithValue(ctx, "id5", 1)
	ctx = ToContext(ctx, factory)
	ctx = context.WithValue(ctx, "id6", 1)
	ctx = context.WithValue(ctx, "id7", 1)
	ctx = context.WithValue(ctx, "id8", 1)
	ctx = context.WithValue(ctx, "id9", 1)
	ctx = context.WithValue(ctx, "id10", 1)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			A(ctx)
		}
	})
}