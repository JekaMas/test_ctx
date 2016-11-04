package product

import (
	"test_ctx/service"
	"test_ctx/service/base"
)

type Product struct {
	base.BaseService
}

// Method which call another service
func (this *Product) ResizeAllImages() int {
	imageService := service.Image(this.Ctx)
	return imageService.Resize()
}

// Independent Method
func (this *Product) AttachImage() int {
	return 10 + 10
}