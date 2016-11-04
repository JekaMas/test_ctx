package main

import (
	"test_ctx/service"
	"context"
	"fmt"
	"test_ctx/common/system"
)

func main() {
	ctx := system.SetupServices(context.Background())

	fmt.Printf("result: %d\n", service.Product(ctx).ResizeAllImages())
	fmt.Printf("result: %d\n", service.Image(ctx).AttachToProduct())
}