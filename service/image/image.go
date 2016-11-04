package image

import (
	"test_ctx/service/base"
	"test_ctx/service"
)

type Image struct{
	base.BaseService
}

// Method which call another service
func (this *Image) AttachToProduct() int {
	productService := service.Product(this.Ctx)
	return productService.AttachImage()
}

// Independent Method
func (this *Image) Resize() int {
	return 1+1
}

