package service

import (
	"context"
	"reflect"
)

type IContextAggregate interface {
	SetContext(ctx context.Context)
}

type IA interface {
	DoA() int
}

type IB interface {
	DoB() int
}

type Factory struct {
	A IA
	B IB
}

var KEY_FACTORY interface{} = "key_factory"

func copy(ctx context.Context, from interface{}) interface{} {
	val := reflect.ValueOf(from)
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}
	result := reflect.New(val.Type()).Interface()
	result.(IContextAggregate).SetContext(ctx)
	return result
}

func AFromContext(ctx context.Context) IA {
	return copy(ctx, ctx.Value(KEY_FACTORY).(*Factory).A).(IA)
}

func BFromContext(ctx context.Context) IB {
	return copy(ctx, ctx.Value(KEY_FACTORY).(*Factory).B).(IB)
}
