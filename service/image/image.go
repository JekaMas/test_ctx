package image

import (
	"test_ctx/service/base"
	"test_ctx/service"
)

type Image struct{
	base.BaseService
}

func (this *Image) DoSomethingWithA() int {
	return service.Product(this.Ctx).DoA()
}

func (this *Image) DoB() int {
	return 1+1
}

