package service

import "context"

//go:generate mockgen -package mocks -destination mocks/a.go test_ctx/service IA
type IA interface {
	IContextAggregate
	DoA() int
	DoSomethingWithB() int
}

func A(ctx context.Context) IA {
	return FromContext(ctx).IA(ctx)
}
