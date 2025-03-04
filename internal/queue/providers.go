package queue

func init() {
	RegisterProvider("SQS", func(config map[string]string) (QueueClient, error) {
		region := config["region"]
		queueURL := config["queueURL"]
		return NewSQSClient(region, queueURL)
	})

	RegisterProvider("RabbitMQ", func(config map[string]string) (QueueClient, error) {
		uri := config["uri"]
		queueName := config["queueName"]
		return NewRabbitMQClient(uri, queueName)
	})
}
