package product_test

import (
	"context"
	"testing"
	"test_ctx/service"
	"github.com/golang/mock/gomock"
	"strconv"
	"test_ctx/common/system"
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

	ctx := system.SetupServices(context.Background())

	mockB := service.MockImage(ctx, mockCtrl)
	mockB.EXPECT().Resize().Return(9998)

	result := service.Product(ctx).ResizeAllImages()
	if result != 9998 {
		t.Fail()
	}
}
