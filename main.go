package main

import (
	"test_ctx/service"
	"context"
	"fmt"
	"test_ctx/common"
)

func main() {
	ctx := common.SetupServices(context.Background())

	fmt.Printf("result: %d\n", service.Product(ctx).DoSomethingWithB())
	fmt.Printf("result: %d\n", service.B(ctx).DoSomethingWithA())
}