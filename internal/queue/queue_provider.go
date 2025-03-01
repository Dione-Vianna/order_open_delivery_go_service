package queue

import (
	"fmt"
)

type QueueProvider string

const (
	SQSProvider       QueueProvider = "SQS"
	RabbitMQProvider  QueueProvider = "RabbitMQ"
)

func NewQueueClient(provider QueueProvider, config map[string]string) (QueueClient, error) {
	switch provider {
	case SQSProvider:
		region := config["region"]
		queueURL := config["queueURL"]
		return NewSQSClient(region, queueURL)
	case RabbitMQProvider:
		uri := config["uri"]
		queueName := config["queueName"]
		return NewRabbitMQClient(uri, queueName)
	default:
		return nil, fmt.Errorf("provvedor de fila desconhecido: %v", provider)
	}
}
