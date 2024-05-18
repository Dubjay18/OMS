package main

import (
	common "github.com/Dubjay18/OMS-common.git"
	pb "github.com/Dubjay18/OMS-common.git/api"
	"net/http"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	CustomerID := r.PathValue("customerID")

	var items []*pb.ItemsWithQuantity
	if err := common.ReadJson(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: CustomerID,
		Items:      items,
	})
}
