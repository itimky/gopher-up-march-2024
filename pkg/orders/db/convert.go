package ordersdb

import (
	contractsdb "github.com/itimky/gopher-up-march-2024/pkg/contracts/db"
	ordersdomain "github.com/itimky/gopher-up-march-2024/pkg/orders/domain"
)

func convertAddOrderParamsToOrder(
	params ordersdomain.AddOrderParams,
) contractsdb.Order {
	return contractsdb.Order{
		UserID:    params.UserID,
		ProductID: params.ProductID,
		OrderID:   params.OrderID,
		Quantity:  params.Quantity,
	}
}
