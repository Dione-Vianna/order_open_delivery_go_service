package queue

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQClient(uri, queueName string) (*RabbitMQClient, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao RabbitMQ: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("erro ao criar canal no RabbitMQ: %v", err)
	}

	queue, err := channel.QueueDeclare(
		queueName, 
		true,      
		false,     
		false,     
		false,     
		nil,       
	)
	if err != nil {
		channel.Close()
		conn.Close()
		return nil, fmt.Errorf("erro ao declarar fila: %v", err)
	}

	return &RabbitMQClient{
		conn:    conn,
		channel: channel,
		queue:   queue,
	}, nil
}

func (client *RabbitMQClient) SendMessage(message string) error {
	err := client.channel.Publish(
		"",               
		client.queue.Name, 
		false,            
		false,            
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Printf("Erro ao enviar mensagem para RabbitMQ: %v", err)
		return fmt.Errorf("erro ao enviar mensagem: %v", err)
	}
	log.Printf("Mensagem enviada para RabbitMQ: %s", message)
	return nil
}
