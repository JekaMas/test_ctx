package service

import (
	"context"
	"reflect"
	"strings"
)

var KEY_SERVICE_FACTORY interface{} = "key_service_factory"

type Factory struct {
	IA
	IB
}

type IContextAggregate interface {
	SetContext(ctx context.Context)
}

func ToContext(ctx context.Context, factory *Factory) context.Context {
	return context.WithValue(ctx, KEY_SERVICE_FACTORY, factory)
}

func FromContext(ctx context.Context) *Factory {
	return ctx.Value(KEY_SERVICE_FACTORY).(*Factory)
}

func CloneService(ctx context.Context, from interface{}) interface{} {
	val := reflect.ValueOf(from)
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

func SetupServices(ctx context.Context) context.Context {
	return ToContext(ctx, &Factory{})
}

