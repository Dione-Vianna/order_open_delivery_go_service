package main

import (
	"log"
	"net"

	proto "order_open_delivery_go_service/api/order-proto"
	"order_open_delivery_go_service/internal/order/handler"
	"order_open_delivery_go_service/internal/order/repository"
	"order_open_delivery_go_service/internal/order/service"
	"order_open_delivery_go_service/internal/queue"

	"google.golang.org/grpc"
)

func main() {

	orderRepository := repository.NewInMemoryOrderRepository()
	sqsClient, err := queue.NewSQSClient("us-east-1", "https://sqs.us-east-1.amazonaws.com/123456789012/myqueue")
	if err != nil {
		log.Fatalf("Erro ao criar cliente SQS: %v", err)
	}

	orderService := service.NewOrderService(orderRepository, sqsClient)

	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

	orderHandler := handler.NewOrderHandler(orderService)

	grpcServer := grpc.NewServer()
	proto.RegisterOrderServiceServer(grpcServer, orderHandler) 

	log.Printf("Servidor gRPC rodando na porta :7777")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}