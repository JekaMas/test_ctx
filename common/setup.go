package common

import (
	"context"
	"test_ctx/service"
	"test_ctx/service/a"
	"test_ctx/service/b"
)

func SetupServices(ctx context.Context) context.Context {
	return service.ToContext(ctx, &service.Factory{
		IA: &a.A{},
		IB: &b.B{},
	})
}
