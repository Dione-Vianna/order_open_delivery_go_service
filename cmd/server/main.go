package main

import (
	"log"
	"net"
	"os"

	proto "order_open_delivery_go_service/api/order-proto"
	"order_open_delivery_go_service/internal/order/handler"
	"order_open_delivery_go_service/internal/order/repository"
	"order_open_delivery_go_service/internal/order/service"
	"order_open_delivery_go_service/internal/queue"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	awsRegion := os.Getenv("AWS_REGION")
	queueURL := os.Getenv("SQS_QUEUE_URL")

	
	sqsClient, err := queue.NewSQSClient(awsRegion, queueURL)

	orderRepository := repository.NewInMemoryOrderRepository()
	
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