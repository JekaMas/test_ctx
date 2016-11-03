package base

import "context"

type BaseService struct {
	Ctx context.Context
}

func (this *BaseService) SetContext(ctx context.Context) {
	this.Ctx = ctx
}


