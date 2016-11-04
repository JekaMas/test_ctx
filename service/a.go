package service

//go:generate mockgen -package mocks -destination mocks/a_gen.go test_ctx/service IA
type IA interface {
	IContextAggregate
	DoA() int
	DoSomethingWithB() int
}

