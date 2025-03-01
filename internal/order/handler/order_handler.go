package handler

import (
	"context"
	"log"

	proto "order_open_delivery_go_service/api/order-proto"
	"order_open_delivery_go_service/internal/order/service"
	"order_open_delivery_go_service/internal/order/validation"
)


type OrderHandler struct {
	proto.UnimplementedOrderServiceServer
	orderService *service.OrderService 
}


func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}


func (handler *OrderHandler) CreateOrder(ctx context.Context, req *proto.OrderRequest) (*proto.OrderResponse, error) {

	log.Println("Chamando CreateOrder no OrderHandler")
	if err := validation.ValidateOrderRequest(req); err != nil {
		log.Printf("Erro de validação: %v", err)
		return nil, err
	}

	
	orderResponse, err := handler.orderService.CreateOrder(ctx, req)
	if err != nil {
		log.Printf("Erro ao criar o pedido: %v", err)
		return nil, err
	}

	log.Printf("Resposta do pedido: %+v", orderResponse)
	
	return orderResponse, nil
}


