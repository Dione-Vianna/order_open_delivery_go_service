package queue

type QueueClient interface {
	SendMessage(message string) error
}

type QueueProvider string
