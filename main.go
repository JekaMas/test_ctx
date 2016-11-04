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
		IA: &a.A{},
		IB: &b.B{},
	}

	ctx := service.ToContext(context.Background(), factory)

	fmt.Printf("result: %d\n", service.A(ctx).DoSomethingWithB())
	fmt.Printf("result: %d\n", service.B(ctx).DoSomethingWithA())
}