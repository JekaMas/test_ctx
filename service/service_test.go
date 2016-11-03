package service

import (
	"testing"
	"context"
)

type AMock struct {
	Ctx context.Context
}

func (this *AMock) SetContext(ctx context.Context) {
	this.Ctx = ctx
}

func (this *AMock) DoA() {}

func BenchmarkAFromContext(b *testing.B) {
	factory := &Factory{
		A: &AMock{},
	}
	ctx := context.WithValue(context.Background(), KEY_FACTORY, factory)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			AFromContext(ctx)
		}
	})
}