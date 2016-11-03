package a

import (
	"context"
	"bigcore/test/service"
	"testing"
	"fmt"
)

type BMock struct {
	Ctx context.Context
}

func (this *BMock) SetContext(ctx context.Context) {
	this.Ctx = ctx
}

func (this *BMock) DoB() int {
	return 2
}

func TestA_DoA(t *testing.T) {
	factory := &service.Factory{B: &BMock{}}
	ctx := context.WithValue(context.Background(), service.KEY_FACTORY, factory)

	result := (&A{ctx}).DoA()
	fmt.Printf("result: %d\n", result)
}