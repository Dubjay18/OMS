package main

import "context"

type Service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *Service {
	return &Service{store: store}
}
func (s *Service) CreateOrder(context.Context) error {
	return nil
}
