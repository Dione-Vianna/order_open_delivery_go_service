package service

import (
	proto "order_open_delivery_go_service/api/order-proto"
	"order_open_delivery_go_service/internal/order/repository"

	"github.com/go-playground/validator"
)



type orderService struct {
	proto.UnimplementedOrderServiceServer
	validator *validator.Validate
	repository repository.OrderRepository
	// queueClient QueueClient  // I need to see how to implement this queue part
}

func NewOrderService(repository repository.OrderRepository, queueClient QueueClient) *orderService {
	return &orderService{
		validator:  validator.New(),
		repository: repository,
		// queueClient: queueClient,
	}
}
