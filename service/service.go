package service

import (
	"context"
	"reflect"
	"strings"
)

var KEY_SERVICE_FACTORY interface{} = "key_service_factory"

type Factory struct {
	IProduct
	IImage
}

type IContextAggregate interface {
	SetContext(ctx context.Context)
}

func ToContext(ctx context.Context, factory *Factory) context.Context {
	return context.WithValue(ctx, KEY_SERVICE_FACTORY, factory)
}

func fromContext(ctx context.Context) *Factory {
	factory :=  ctx.Value(KEY_SERVICE_FACTORY)
	if factory == nil {
		panic("You forgot to call common.SetupServices")
	}
	return factory.(*Factory)
}

func cloneService(ctx context.Context, from interface{}) interface{} {
	val := reflect.ValueOf(from)

	// TODO: remove from production code
	if strings.Contains(val.Type().String(), "mock") {
		return from
	}
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}
	result := reflect.New(val.Type()).Interface()
	result.(IContextAggregate).SetContext(ctx)
	return result
}


