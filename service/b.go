package service

import (
	"context"
)

//go:generate mockgen -package mocks -destination mocks/b.go test_ctx/service IB
type IB interface {
	IContextAggregate
	DoB() int
	DoSomethingWithA() int
}

func B(ctx context.Context) IB {
	return FromContext(ctx).IB(ctx)
}
