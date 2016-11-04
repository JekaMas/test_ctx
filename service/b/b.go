package b

import (
	"test_ctx/service/base"
	"test_ctx/service"
)

type B struct{
	base.BaseService
}

func (this *B) DoSomethingWithA() int {
	return service.A(this.Ctx).DoA()
}

func (this *B) DoB() int {
	return 1+1
}

