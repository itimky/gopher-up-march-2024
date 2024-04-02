package main

import (
	ordersdb "github.com/itimky/gopher-up-march-2024/pkg/orders/db"
	ordersdomain "github.com/itimky/gopher-up-march-2024/pkg/orders/domain"
	ordershttp "github.com/itimky/gopher-up-march-2024/pkg/orders/http"
)

func main() {
	db := ordersdb.NewInMemDB()
	orders := ordersdomain.NewOrders(db)
	httpHandler := ordershttp.NewHandler(orders)
	_ = httpHandler
}
