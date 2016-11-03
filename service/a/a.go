package a

import (
	"context"
	"bigcore/test/service"
)

type A struct {
	Ctx context.Context
}

func (this *A) SetContext(ctx context.Context) {
	this.Ctx = ctx
}

func (this *A) DoA() int {
	return service.BFromContext(this.Ctx).DoB()
}