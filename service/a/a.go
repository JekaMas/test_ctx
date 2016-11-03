package a

import (
	"bigcore/test/service"
	"test_ctx/service/base"
)

type A struct {
	base.BaseService
}

func (this *A) DoA() int {
	return service.B(this.Ctx).DoB()
}