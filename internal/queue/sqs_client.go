package queue

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)


type SQSClient struct {
    sqsService *sqs.SQS
    queueURL   string
}

func NewSQSClient(region, queueURL string) (*SQSClient, error) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao criar a sessão do AWS: %v", err)
	}

	svc := sqs.New(sess)

	return &SQSClient{
		sqsService: svc,
		queueURL:   queueURL,
	}, nil
}


func (client *SQSClient) SendMessage(message string) error {

	_, err := client.sqsService.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    &client.queueURL,
		MessageBody: aws.String(message),
	})
	if err != nil {
		log.Printf("Erro ao enviar mensagem para SQS: %v", err)
		return fmt.Errorf("erro ao enviar mensagem: %v", err)
	}
	log.Printf("Mensagem enviada para SQS: %s", message)
	return nil
}
