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

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		panic("Error loadin .env file")
	}

	provider := queue.QueueProvider(os.Getenv("QUEUE_PROVIDER"))

	config := make(map[string]string)

	switch provider {
	case "SQS":
		config["region"] = os.Getenv("AWS_REGION")
		config["queueURL"] = os.Getenv("SQS_QUEUE_URL")
	case "RabbitMQ":
		config["uri"] = os.Getenv("RABBITMQ_URI")
		config["queueName"] = os.Getenv("RABBITMQ_QUEUE_NAME")
	case "Kafka":
		config["region"] = os.Getenv("AWS_REGION")
		config["queueURL"] = os.Getenv("KAFKA_QUEUE_URL")

	default:
		log.Fatalf("Provedor de fila desconhecido: %s", provider)
	}

	if err := startServer(provider, config); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
