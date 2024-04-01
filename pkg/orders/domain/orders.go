package ordersdomain

import (
	"context"
	"fmt"
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
	orderID := GenerateRandomID(ctx)

	err := o.db.AddOrder(ctx, convertCreateOrderParamsToAddOrderParams(params, orderID))
	if err != nil {
		return nil, fmt.Errorf("add order: %w", err)
	}

	return &CreateOrderResult{
		OrderID: orderID,
	}, nil
}
