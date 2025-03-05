package main

import (
	"log"
	"os"
	"testing"
	"time"

	"order_open_delivery_go_service/internal/queue"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestQueueProviders(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Error loadin .env file")
	}

	config := map[string]string{
		"region":   os.Getenv("AWS_REGION_TEST"),
		"queueURL": os.Getenv("SQS_QUEUE_URL_TEST"),
	}

	log.Printf("environment variables %v", config)

	errChan := make(chan error, 1)

	go func() {
		errChan <- startServer("SQS", config)
	}()

	time.Sleep(2 * time.Second)

	select {
	case err := <-errChan:
		if err != nil {
			t.Fatalf("Erro ao iniciar o servidor: %v", err)
		}
	default:

	}
	client, err := queue.NewQueueClient("SQS", config)
	assert.NoError(t, err, "Error creating queue client")

	err = client.SendMessage("Test message")
	assert.NoError(t, err, "Error sending message")

}
