package system

import (
	"context"
	"test_ctx/service"
	"test_ctx/service/product"
	"test_ctx/service/image"
)

func SetupServices(ctx context.Context) context.Context {
	return service.ToContext(ctx, &service.Factory{
		IProduct: &product.Product{},
		IImage: &image.Image{},
	})
}
