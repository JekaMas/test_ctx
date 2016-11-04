package main

import (
	"test_ctx/service"
	"context"
	"fmt"
	"test_ctx/common/system"
)

func main() {
	ctx := system.SetupServices(context.Background())

	// i was lazy to create handlers with context
	// imagine that now you are in handler and already have context.
	// let's create and call couple of services

	resizeResult := service.Product(ctx).ResizeAllImages()
	fmt.Printf("result: %d\n", resizeResult)

	attachResult := service.Image(ctx).AttachToProduct()
	fmt.Printf("result: %d\n", attachResult)
}