package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartServer(t *testing.T) {

	config := make(map[string]string)

	config["region"] = os.Getenv("AWS_REGION_TEST")
	config["queueURL"] = os.Getenv("SQS_QUEUE_URL_TEST")

	err := startServer("SQS", config)

	if err != nil {
		assert.EqualError(t, err, err.Error())
	} else {
		assert.NoError(t, err)
	}

}
