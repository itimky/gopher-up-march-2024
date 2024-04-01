package ordersdomain

import "github.com/itimky/gopher-up-march-2024/pkg"

type CreateOrderParams struct {
	ProductID pkg.ProductID
	UserID    pkg.UserID
	Quantity  int
}

type CreateOrderResult struct {
	OrderID pkg.OrderID
}

type AddOrderParams struct {
	ProductID pkg.ProductID
	UserID    pkg.UserID
	OrderID   pkg.OrderID
	Quantity  int
}
