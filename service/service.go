package service

import (
	"context"
	"reflect"
)

type IContextAggregate interface {
	SetContext(ctx context.Context)
}

type Factory struct {
	A IA
	B IB
}

var KEY_FACTORY interface{} = "key_factory"

func ToContext(ctx context.Context, factory *Factory) context.Context {
	return context.WithValue(ctx, KEY_FACTORY, factory)
}

func FromContext(ctx context.Context) *Factory {
	return ctx.Value(KEY_FACTORY).(*Factory)
}

func clone(ctx context.Context, from interface{}) interface{} {
	val := reflect.ValueOf(from)
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}
	result := reflect.New(val.Type()).Interface()
	result.(IContextAggregate).SetContext(ctx)
	return result
}

