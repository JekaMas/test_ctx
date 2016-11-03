package b

import (
	"context"
)

type B struct{
	Ctx context.Context
}

func (this *B) SetContext(ctx context.Context) {
	this.Ctx = ctx
}

func (*B) DoB() int {
	return 1
}