package queue

import (
	"fmt"
	"sync"
)

type QueueClientFactory func(config map[string]string) (QueueClient, error)

var (
	providers   = make(map[QueueProvider]QueueClientFactory)
	providersMu sync.RWMutex
)

func RegisterProvider(name QueueProvider, factory QueueClientFactory) {
	providersMu.Lock()
	defer providersMu.Unlock()
	providers[name] = factory
}

func NewQueueClient(provider QueueProvider, config map[string]string) (QueueClient, error) {
	providersMu.RLock()
	defer providersMu.RUnlock()

	factory, exists := providers[provider]
	if !exists {
		return nil, fmt.Errorf("provedor de fila desconhecido: %v", provider)
	}

	return factory(config)
}
