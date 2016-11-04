package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"test_ctx/service/mocks"
)


type IFactory interface {
	IProduct(ctx context.Context) IProduct
	IB(ctx context.Context) IImage
}


// internal methods
func (this *Factory) Product(ctx context.Context) IProduct {
	return cloneService(ctx, this.IProduct).(IProduct)
}

func (this *Factory) B(ctx context.Context) IImage {
	return cloneService(ctx, this.IImage).(IImage)
}


// Getters
func Product(ctx context.Context) IProduct {
	return fromContext(ctx).Product(ctx)
}

func B(ctx context.Context) IImage {
	return fromContext(ctx).B(ctx)
}


// Mocks
func MockProduct(ctx context.Context, ctrl *gomock.Controller) *mocks.MockIProduct {
	mock := mocks.NewMockIProduct(ctrl)
	fromContext(ctx).IProduct = mock
	return mock
}

func MockImage(ctx context.Context, ctrl *gomock.Controller) *mocks.MockIImage {
	mock := mocks.NewMockIImage(ctrl)
	fromContext(ctx).IImage = mock
	return mock
}

