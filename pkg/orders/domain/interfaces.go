package ordersdomain

import "context"

type db interface {
	AddOrder(ctx context.Context, params AddOrderParams) error
}
