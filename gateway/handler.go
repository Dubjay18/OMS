package main

import (
	"errors"
	common "github.com/Dubjay18/OMS-common.git"
	pb "github.com/Dubjay18/OMS-common.git/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	if err := validateItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err)
		return
	}
	o, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: CustomerID,
		Items:      items,
	})

	rStatus := status.Convert(err)

	if rStatus != nil {
		if rStatus.Code() == codes.InvalidArgument {
			common.WriteError(w, http.StatusBadRequest, rStatus.Err())
			return
		}
	}
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err)
		return

	}
	common.WriteJson(w, http.StatusCreated, o)
}

func validateItems(items []*pb.ItemsWithQuantity) error {
	if len(items) == 0 {
		return common.ErrorNoItems
	}
	for _, item := range items {
		if item.ItemID == "" {
			return errors.New("item_id is required")
		}
		if item.Quantity <= 0 {
			return errors.New("quantity must be greater than 0")
		}
	}
	return nil
}
