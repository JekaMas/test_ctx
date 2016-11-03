package a

import (
	"context"
	"testing"
	"fmt"
	"test_ctx/service/base"
	"test_ctx/service"
	"github.com/golang/mock/gomock"
	"test_ctx/service/mocks"
)

type BMock struct {
	base.BaseService
}

func (this *BMock) DoB() int {
	return 9999
}

func (this *BMock) DoSomethingWithA() int {
	return 0
}

func newA(ctx context.Context) *A {
	return &A{base.BaseService{ctx}}
}

func TestA_DoA(t *testing.T) {
	factory := &service.Factory{B: &BMock{}}
	ctx := service.ToContext(context.Background(), factory)

	result := newA(ctx).DoSomethingWithB()
	fmt.Printf("result: %d\n", result)
}

func TestA_DoA2(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockB := mocks.NewMockIB(mockCtrl)
	mockB.EXPECT().DoB().Return(9998)

	factory := &service.TestFactory{B: mockB}
	ctx := service.ToContext(context.Background(), factory)

	result := newA(ctx).DoSomethingWithB()
	fmt.Printf("result: %d\n", result)
}