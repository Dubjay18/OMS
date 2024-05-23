package main

import (
	"context"
	common "github.com/Dubjay18/OMS-common.git"
	pb "github.com/Dubjay18/OMS-common.git/api"
	"log"
)

type Service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *Service {
	return &Service{store: store}
}
func (s *Service) CreateOrder(context.Context) error {
	return nil
}

func (s *Service) ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) error {
	if len(p.Items) == 0 {
		return common.ErrorNoItems
	}
	mergedItems := mergeItemQuantities(p.Items)
	log.Print(mergedItems)
	//TODO: Validate items WITH STOCK SERVICE
	return nil
}

func mergeItemQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
	mergedItems := make([]*pb.ItemsWithQuantity, 0)
	for _, item := range items {
		found := false
		for _, mergedItem := range mergedItems {
			if item.ItemID == mergedItem.ItemID {
				mergedItem.Quantity += item.Quantity
				found = true
				break
			}
		}
		if !found {
			mergedItems = append(mergedItems, item)
		}
	}
	return mergedItems
}
