package product

import (
	"test_ctx/service"
	"test_ctx/service/base"
)

type Product struct {
	base.BaseService
}

func (this *Product) ResizeAllImages() int {
	imageService := service.Image(this.Ctx)
	return imageService.Resize()
}

func (this *Product) AttachImage() int {
	return 10 + 10
}