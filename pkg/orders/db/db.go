package ordersdb

import (
	"context"

	contractsdb "github.com/itimky/gopher-up-march-2024/pkg/contracts/db"
	ordersdomain "github.com/itimky/gopher-up-march-2024/pkg/orders/domain"
)

type InMemDB struct {
	orders []contractsdb.Order
}

func NewInMemDB() *InMemDB {
	return &InMemDB{
		orders: nil,
	}
}

func (d *InMemDB) AddOrder(
	_ context.Context,
	params ordersdomain.AddOrderParams,
) error {
	d.orders = append(d.orders, convertAddOrderParamsToOrder(params))

	return nil
}
