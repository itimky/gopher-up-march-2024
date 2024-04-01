package ordershttp

import (
	"context"

	ordersdomain "github.com/itimky/gopher-up-march-2024/pkg/orders/domain"
)

type orders interface {
	CreateOrder(ctx context.Context, params ordersdomain.CreateOrderParams) (*ordersdomain.CreateOrderResult, error)
}
