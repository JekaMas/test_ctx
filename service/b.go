package service

//go:generate mockgen -package mocks -destination mocks/b_gen.go test_ctx/service IB
type IB interface {
	IContextAggregate
	DoB() int
	DoSomethingWithA() int
}


