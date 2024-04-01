package ordersdomain

import "github.com/itimky/gopher-up-march-2024/pkg"

func convertCreateOrderParamsToAddOrderParams(
	createOrderParams CreateOrderParams,
	orderID pkg.OrderID,
) AddOrderParams {
	return AddOrderParams{
		UserID:    createOrderParams.UserID,
		ProductID: createOrderParams.ProductID,
		OrderID:   orderID,
		Quantity:  createOrderParams.Quantity,
	}
}
