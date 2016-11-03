package service

import (
	"context"
	"reflect"
)

var KEY_SERVICE_FACTORY interface{} = "key_service_factory"

type IContextAggregate interface {
	SetContext(ctx context.Context)
}

type Factory struct {
	A IA
	B IB
}

type IFactory interface {
	IA(ctx context.Context) IA
	IB(ctx context.Context) IB
}

func (this *Factory) IA(ctx context.Context) IA {
	return this.CloneService(ctx, this.A).(IA)
}

func (this *Factory) IB(ctx context.Context) IB {
	return this.CloneService(ctx, this.B).(IB)
}

type TestFactory struct {
	A IA
	B IB
}

func (this *TestFactory) IA(ctx context.Context) IA {
	return this.A
}

func (this *TestFactory) IB(ctx context.Context) IB {
	return this.B
}

func ToContext(ctx context.Context, factory IFactory) context.Context {
	return context.WithValue(ctx, KEY_SERVICE_FACTORY, factory)
}

func FromContext(ctx context.Context) IFactory {
	return ctx.Value(KEY_SERVICE_FACTORY).(IFactory)
}

func (*Factory) CloneService(ctx context.Context, from interface{}) interface{} {
	val := reflect.ValueOf(from)
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}
	result := reflect.New(val.Type()).Interface()
	result.(IContextAggregate).SetContext(ctx)
	return result
}

