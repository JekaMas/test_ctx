package service

import "context"

//go:generate mockgen -package mocks -destination mocks/a.go test_ctx/service IA
type IA interface {
	IContextAggregate
	DoA() int
	DoSomethingWithB() int
}

func A(ctx context.Context) IA {
	return clone(ctx, FromContext(ctx).A).(IA)
}
