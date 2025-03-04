package queue

type QueueClient interface {
	SendMessage(message string) error
}

type QueueProvider string

type QueueClientFactory func(config map[string]string) (QueueClient, error)
