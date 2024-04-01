package contractsdb

import "github.com/itimky/gopher-up-march-2024/pkg"

type Order struct {
	OrderID   pkg.OrderID
	ProductID pkg.ProductID
	UserID    pkg.UserID
	Quantity  int
}
