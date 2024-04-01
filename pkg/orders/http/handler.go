package ordershttp

import (
	"encoding/json"
	"net/http"

	contractshttp "github.com/itimky/gopher-up-march-2024/pkg/contracts/http"
)

type Handler struct {
	orders orders
}

func NewHandler(
	orders orders,
) *Handler {
	return &Handler{
		orders: orders,
	}
}

func (h *Handler) Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/orders/v1", h.CreateOrderV1)
	return mux
}

func (h *Handler) CreateOrderV1(
	w http.ResponseWriter,
	r *http.Request,
) {
	var orderV1 contractshttp.CreateOrderRequestV1
	if err := json.NewDecoder(r.Body).Decode(&orderV1); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.orders.CreateOrder(r.Context(), convertCreateOrderRequestV1ToCreateOrderParams(orderV1))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(convertCreateOrderResultToCreateOrderResponseV1(*result)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
