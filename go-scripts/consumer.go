package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    start := time.Now()

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

	if err != nil {
		panic(err)
	}

    messagesPerBatch := 1_000_000
    nReceived := 0

    for {
        msg, err := c.ReadMessage(time.Second)
        if err == nil {
            nReceived++
        } else if !err.(kafka.Error).IsTimeout() {
            // Timeout is not considered an error because it is raised by
            // ReadMessage in absence of messages.
            fmt.Printf("Consumer error: %v (%v)\n", err, msg)
        }

        if (nReceived % messagesPerBatch == 0) && (nReceived > 0) {
            elapsed := time.Since(start)
            fmt.Printf("Received %d messages within %.2f seconds\n", nReceived, elapsed.Seconds())
        }
    }

	c.Close()
}