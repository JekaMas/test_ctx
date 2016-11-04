package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"test_ctx/service/mocks"
)

type IFactory interface {
	IProduct(ctx context.Context) IProduct
	IImage(ctx context.Context) IImage
}

// Getters
func Product(ctx context.Context) IProduct {
	return cloneService(ctx, fromContext(ctx).IProduct).(IProduct)
}

func Image(ctx context.Context) IImage {
	return cloneService(ctx, fromContext(ctx).IImage).(IImage)
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
