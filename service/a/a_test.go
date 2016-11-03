package a

import (
	"context"
	"bigcore/test/service"
	"testing"
	"fmt"
	"test_ctx/service/base"
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
	ctx := context.WithValue(context.Background(), service.KEY_FACTORY, factory)

	result := newA(ctx).DoA()
	fmt.Printf("result: %d\n", result)
}