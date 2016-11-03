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

	ctx := service.ToContext(context.Background(), factory)

	result := service.A(ctx).DoA()
	fmt.Printf("result: %d\n", result)
}