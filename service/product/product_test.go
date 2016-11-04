package product_test

import (
	"context"
	"testing"
	"test_ctx/service"
	"github.com/golang/mock/gomock"
	"strconv"
	"test_ctx/common"
)

func TestParallel(t *testing.T) {
	for i := 0; i < 10000; i++ {
		t.Run("Tet" + strconv.Itoa(i), parallel)
	}
}

func parallel(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := common.SetupServices(context.Background())

	mockB := service.MockImage(ctx, mockCtrl)
	mockB.EXPECT().DoB().Return(9998)

	result := service.Product(ctx).DoSomethingWithB()
	if result != 9998 {
		t.Fail()
	}
}
