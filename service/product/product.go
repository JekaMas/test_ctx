package product

import (
	"test_ctx/service"
	"test_ctx/service/base"
)

type Product struct {
	base.BaseService
}

func (this *Product) ResizeAllImages() int {
	return service.Image(this.Ctx).Resize()
}

func (this *Product) AttachImage() int {
	return 10
}