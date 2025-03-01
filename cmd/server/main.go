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
		log.Fatalf("Erro ao criar cliente de fila: %v", err)
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

	log.Printf("Servidor gRPC rodando na porta :7777")
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	provider := queue.SQSProvider // RabbitMQProvider or SQSProvider

	config := make(map[string]string)

	if provider == queue.SQSProvider {
		config["region"] = os.Getenv("AWS_REGION")
		config["queueURL"] = os.Getenv("SQS_QUEUE_URL")
	}

	if provider == queue.RabbitMQProvider {
		config["uri"] = os.Getenv("RABBITMQ_URI")
		config["queueName"] = os.Getenv("RABBITMQ_QUEUE_NAME")
	}

	if err := startServer(provider, config); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
