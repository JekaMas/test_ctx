package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"test_ctx/service/mocks"
)

// generate me
type IFactory interface {
	IA(ctx context.Context) IA
	IB(ctx context.Context) IB
}

// internal methods
func (this *Factory) A(ctx context.Context) IA {
	return CloneService(ctx, this.IA).(IA)
}

func (this *Factory) B(ctx context.Context) IB {
	return CloneService(ctx, this.IB).(IB)
}


// Getters
func A(ctx context.Context) IA {
	return FromContext(ctx).A(ctx)
}

func B(ctx context.Context) IB {
	return FromContext(ctx).B(ctx)
}


// Mocks
func MockA(ctx context.Context, ctrl *gomock.Controller) *mocks.MockIA {
	mock := mocks.NewMockIA(ctrl)
	FromContext(ctx).IA = mock
	return mock
}

func MockB(ctx context.Context, ctrl *gomock.Controller) *mocks.MockIB {
	mock := mocks.NewMockIB(ctrl)
	FromContext(ctx).IB = mock
	return mock
}

