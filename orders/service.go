package main

import "context"

type Service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *Service {
	return &Service{store: store}
}
func CreateOrder(ctx context.Context) error {
	return nil
}
