package a

import (
	"context"
	"testing"
	"test_ctx/service/base"
	"test_ctx/service"
	"github.com/golang/mock/gomock"
	"strconv"
)

func newA(ctx context.Context) *A {
	return &A{base.BaseService{ctx}}
}

func TestParallel(t *testing.T) {
	for i := 0; i < 10000; i++ {
		t.Run("Tet" + strconv.Itoa(i), parallel)
	}
}

func parallel(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := service.SetupServices(context.Background())
	mockB := service.MockB(ctx, mockCtrl)
	mockB.EXPECT().DoB().Return(9998)

	result := newA(ctx).DoSomethingWithB()
	if result != 9998 {
		t.Fail()
	}
}
