package main

import (
	"test_ctx/service"
	"context"
	"test_ctx/service/a"
	"test_ctx/service/b"
	"fmt"
)

func main() {
	factory := &service.Factory{
		A: &a.A{},
		B: &b.B{},
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, service.KEY_FACTORY, factory)

	result := service.A(ctx).DoA()
	fmt.Printf("result: %d\n", result)
}