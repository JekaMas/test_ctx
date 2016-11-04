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
	return CloneService(ctx, this.IProduct).(IProduct)
}

func (this *Factory) B(ctx context.Context) IImage {
	return CloneService(ctx, this.IImage).(IImage)
}


// Getters
func Product(ctx context.Context) IProduct {
	return FromContext(ctx).Product(ctx)
}

func B(ctx context.Context) IImage {
	return FromContext(ctx).B(ctx)
}


// Mocks
func MockProduct(ctx context.Context, ctrl *gomock.Controller) *mocks.MockIProduct {
	mock := mocks.NewMockIProduct(ctrl)
	FromContext(ctx).IProduct = mock
	return mock
}

func MockImage(ctx context.Context, ctrl *gomock.Controller) *mocks.MockIImage {
	mock := mocks.NewMockIImage(ctrl)
	FromContext(ctx).IImage = mock
	return mock
}

