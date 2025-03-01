package repository

import (
	"sync"

	proto "order_open_delivery_go_service/api/order-proto"
)

type OrderRepository interface {
	Save(order *proto.OrderRequest) error
}

type InMemoryOrderRepository struct {
	mu     sync.Mutex
	orders map[string]*proto.OrderRequest
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		orders: make(map[string]*proto.OrderRequest),
	}
}

func (repository *InMemoryOrderRepository) Save(order *proto.OrderRequest) error {
	repository.mu.Lock()
	defer repository.mu.Unlock()
	repository.orders[order.Id] = order
	return nil
}


