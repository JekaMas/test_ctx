package image

import (
	"test_ctx/service/base"
	"test_ctx/service"
)

type Image struct{
	base.BaseService
}

func (this *Image) AttachToProduct() int {
	return service.Product(this.Ctx).AttachImage()
}

func (this *Image) Resize() int {
	return 1+1
}

