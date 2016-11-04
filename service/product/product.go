package product

import (
	"test_ctx/service"
	"test_ctx/service/base"
)

type Product struct {
	base.BaseService
}

func (this *Product) DoSomethingWithB() int {
	return service.B(this.Ctx).DoB()
}

func (this *Product) DoA() int {
	return 10
}