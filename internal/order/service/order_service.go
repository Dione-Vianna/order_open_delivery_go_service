package service

import (
	"context"
	"encoding/json"
	"log"
	proto "order_open_delivery_go_service/api/order-proto"
	"order_open_delivery_go_service/internal/order/repository"
	"order_open_delivery_go_service/internal/queue"

	"github.com/go-playground/validator"
)

type OrderService struct {
	proto.UnimplementedOrderServiceServer
	validator   *validator.Validate
	repository  repository.OrderRepository
	queueClient queue.QueueClient
}

func NewOrderService(repository repository.OrderRepository, queueClient queue.QueueClient) *OrderService {
	return &OrderService{
		validator:   validator.New(),
		repository:  repository,
		queueClient: queueClient,
	}
}

func (service *OrderService) CreateOrder(ctx context.Context, req *proto.OrderRequest) (*proto.OrderResponse, error) {
	log.Println("Chamando CreateOrder no OrderService")

	err := service.repository.Save(req)
	if err != nil {
		log.Printf("Erro ao salvar pedido: %v", err)
		return nil, err
	}

	orderJSON, err := json.Marshal(req)
	if err != nil {
		log.Printf("Erro ao converter pedido para JSON: %v", err)
		return nil, err
	}

	err = service.queueClient.SendMessage(string(orderJSON))
	if err != nil {
		log.Printf("Erro ao enviar mensagem para fila: %v", err)
		return nil, err
	}

	return &proto.OrderResponse{
		Status:  "success",
		Message: "Pedido criado com sucesso!",
	}, nil
}
