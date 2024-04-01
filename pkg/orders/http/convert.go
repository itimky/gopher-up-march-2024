package ordershttp

import (
	"github.com/itimky/gopher-up-march-2024/pkg"
	contractshttp "github.com/itimky/gopher-up-march-2024/pkg/contracts/http"
	ordersdomain "github.com/itimky/gopher-up-march-2024/pkg/orders/domain"
)

func convertCreateOrderRequestV1ToCreateOrderParams(
	createOrderV1 contractshttp.CreateOrderRequestV1,
) ordersdomain.CreateOrderParams {
	return ordersdomain.CreateOrderParams{
		ProductID: pkg.ProductID(createOrderV1.ProductID),
		UserID:    pkg.UserID(createOrderV1.UserID),
		Quantity:  int(createOrderV1.Quantity),
	}
}

func convertCreateOrderResultToCreateOrderResponseV1(
	createOrderResult ordersdomain.CreateOrderResult,
) contractshttp.CreateOrderResponseV1 {
	return contractshttp.CreateOrderResponseV1{
		OrderID: string(createOrderResult.OrderID),
	}
}
