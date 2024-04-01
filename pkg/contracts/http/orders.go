package contractshttp

type CreateOrderRequestV1 struct {
	ProductID string `json:"productId"`
	UserID    string `json:"userId"`
	Quantity  uint16 `json:"quantity"`
}

type CreateOrderResponseV1 struct {
	OrderID string `json:"orderId"`
}
