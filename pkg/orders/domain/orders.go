package ordersdomain

import (
	"context"
	"fmt"

	"github.com/itimky/gopher-up-march-2024/pkg"
	"github.com/itimky/gopher-up-march-2024/pkg/sys"
)

type Orders struct {
	db db
}

func NewOrders(
	db db,
) *Orders {
	return &Orders{
		db: db,
	}
}

func (o *Orders) CreateOrder(
	ctx context.Context,
	params CreateOrderParams,
) (*CreateOrderResult, error) {
	orderID, err := sys.GenerateRandomID(ctx)
	if err != nil {
		return nil, fmt.Errorf("generate random id: %w", err)
	}

	err = o.db.AddOrder(ctx, convertCreateOrderParamsToAddOrderParams(params, pkg.OrderID(orderID)))
	if err != nil {
		return nil, fmt.Errorf("add order: %w", err)
	}

	return &CreateOrderResult{
		OrderID: pkg.OrderID(orderID),
	}, nil
}
