package a

import (
	"context"
	"testing"
	"fmt"
	"test_ctx/service/base"
	"test_ctx/service"
)

type BMock struct {
	base.BaseService
}

func (this *BMock) DoB() int {
	return 2
}

func newA(ctx context.Context) *A {
	return &A{base.BaseService{ctx}}
}

func TestA_DoA(t *testing.T) {
	factory := &service.Factory{B: &BMock{}}
	ctx := service.ToContext(context.Background(), factory)

	result := newA(ctx).DoA()
	fmt.Printf("result: %d\n", result)
}