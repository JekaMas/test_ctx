package service

//go:generate mockgen -package mocks -destination mocks/image_gen.go test_ctx/service IImage
type IImage interface {
	IContextAggregate
	Resize() int
	AttachToProduct() int
}


