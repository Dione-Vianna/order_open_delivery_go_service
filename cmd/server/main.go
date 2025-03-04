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

func startServer(provider queue.QueueProvider, config map[string]string) error {
	client, err := queue.NewQueueClient(provider, config)
	if err != nil {
		log.Fatalf("Error creating queue client: %v", err)
		panic("Error creating queue client")
	}

	orderRepository := repository.NewInMemoryOrderRepository()
	orderService := service.NewOrderService(orderRepository, client)

	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		return err
	}

	orderHandler := handler.NewOrderHandler(orderService)
	grpcServer := grpc.NewServer()
	proto.RegisterOrderServiceServer(grpcServer, orderHandler)

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	log.Printf("gRPC server running on port :7777")
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Error loadin .env file")
	}

	providerStr := os.Getenv("QUEUE_PROVIDER")
	var provider queue.QueueProvider

	switch providerStr {
	case "SQS":
		provider = "SQS"
	case "RabbitMQ":
		provider = "RabbitMQ"
	default:
		log.Fatalf("Provedor de fila desconhecido: %s", providerStr)
	}

	config := make(map[string]string)

	if provider == "SQS" {
		config["region"] = os.Getenv("AWS_REGION")
		config["queueURL"] = os.Getenv("SQS_QUEUE_URL")
	}

	if provider == "RabbitMQ" {
		config["uri"] = os.Getenv("RABBITMQ_URI")
		config["queueName"] = os.Getenv("RABBITMQ_QUEUE_NAME")
	}

	if err := startServer(provider, config); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
