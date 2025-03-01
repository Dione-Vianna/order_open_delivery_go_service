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

	provider := queue.RabbitMQProvider // RabbitMQProvider or SQSProvider

	config := make(map[string]string)

	if provider == queue.SQSProvider {
		config["region"] = os.Getenv("AWS_REGION")
		config["queueURL"] = os.Getenv("SQS_QUEUE_URL")
	}

	if provider == queue.RabbitMQProvider {
		config["uri"] = os.Getenv("RABBITMQ_URI")
		config["queueName"] = os.Getenv("RABBITMQ_QUEUE_NAME")
	}

	sqsClient, err := queue.NewQueueClient(provider, config)
	if err != nil {
		log.Fatalf("Erro ao criar cliente SQS: %v", err)
	}

	orderRepository := repository.NewInMemoryOrderRepository()

	orderService := service.NewOrderService(orderRepository, sqsClient)

	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

	orderHandler := handler.NewOrderHandler(orderService)

	grpcServer := grpc.NewServer()
	proto.RegisterOrderServiceServer(grpcServer, orderHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

	log.Printf("Servidor gRPC rodando na porta :7777")
}
