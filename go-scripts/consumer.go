package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    start := time.Now()

    // https://github.com/confluentinc/librdkafka/blob/master/CONFIGURATION.md
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id": "myGroup",
		"auto.offset.reset": "earliest",
        // this makes a significant difference for consumption speed
        "fetch.message.max.bytes": 1_000_000_000, // 1 GB (largest allowed)
        // this makes a minor difference for the consumption speed
        "fetch.max.bytes": 1_000_000_000,  // 1 GB (~2 GB is largest allowed, 50 MB is default)
        "queued.max.messages.kbytes": 1_000_000, // 1 GB (~2 GB is largest allowed, ~65 MB is default)
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