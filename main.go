package main

import (
	"bigcore/test/service"
	"context"
	"bigcore/test/service/a"
	"bigcore/test/service/b"
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