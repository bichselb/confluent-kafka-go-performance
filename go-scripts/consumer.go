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

	err = c.Subscribe("myTopic", nil)

	if err != nil {
		panic(err)
	}

    messagesPerBatch := 1_000_000
    nReceived := 0

    for {
        msg, err := c.ReadMessage(10 * time.Second)
        if err == nil {
            nReceived++
        } else if err.(kafka.Error).IsTimeout() {
            fmt.Printf("Timeout after %d messages\n", nReceived)
            break
        } else {
            fmt.Printf("Consumer error: %v (%v)\n", err, msg)
        }

        if (nReceived % messagesPerBatch == 0) && (nReceived > 0) {
            elapsed := time.Since(start)
            perSecond := int(float64(nReceived) / elapsed.Seconds())
            fmt.Printf("Received %d messages within %.2f seconds (average %d messages/sec)\n", nReceived, elapsed.Seconds(), perSecond)
        }
    }

	c.Close()
}