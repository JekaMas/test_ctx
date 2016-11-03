package a

import (
	"test_ctx/service"
	"test_ctx/service/base"
)

type A struct {
	base.BaseService
}

func (this *A) DoA() int {
	return service.B(this.Ctx).DoB()
}