package main

type store struct {
	// TODO: add mongo db instance
}

func NewStore(st OrdersStore) *store {
	return &store{
		store: st,
	}
}

func (s *store) CreateOrder() error {
	return nil
}
