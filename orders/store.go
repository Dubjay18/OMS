package main

import "context"

type Store struct {
	// Todo: Add mongodb instance
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Create(ctx context.Context) error {
	return nil
}
