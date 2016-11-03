package main

import (
	"bigcore/test/service"
	"context"
	"bigcore/test/service/a"
	"bigcore/test/service/b"
)

func main() {
	factory := &service.Factory{
		A: &a.A{},
		B: &b.B{},
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, service.KEY_FACTORY, factory)

	service.AFromContext(ctx).DoA()
	service.BFromContext(ctx).DoB()
}