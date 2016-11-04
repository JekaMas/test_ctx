package service

//go:generate mockgen -package mocks -destination mocks/product_gen.go test_ctx/service IProduct
type IProduct interface {
	IContextAggregate
	AttachImage() int
	ResizeAllImages() int
}

