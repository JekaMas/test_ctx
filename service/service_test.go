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

func BenchmarkAFromContext(b *testing.B) {
	factory := &Factory{
		A: &AMock{},
	}
	ctx := context.WithValue(context.Background(), KEY_FACTORY, factory)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			A(ctx)
		}
	})
}