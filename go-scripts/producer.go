package main

import (
    "fmt"
    "strconv"
    "time"

    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    start := time.Now()

    // https://github.com/confluentinc/librdkafka/blob/master/CONFIGURATION.md
    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "kafka:29092",
        // avoid message drops
        "queue.buffering.max.messages": 0,  // Maximum number of messages allowed on the producer queue, 0 disables this limit
        // detailed logging
        // "debug": "msg",
    })
    if err != nil {
        panic(err)
    }
    defer p.Close()

    messagesDelivered := 0

    // Delivery report handler for produced messages
    go func() {
        for e := range p.Events() {
            switch ev := e.(type) {
            case *kafka.Message:
                messagesDelivered ++
                if ev.TopicPartition.Error != nil {
                    fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
                }
            }
        }
    }()

    topic := "myTopic"
    messagesPerBatch := 1_000_000
    nBatches := 30

    for i := 0; i < nBatches; i++ {
        for j := 0; j < messagesPerBatch; j++ {
            message := "Batch " + strconv.Itoa(i) + ", message " + strconv.Itoa(j)
            p.Produce(&kafka.Message{
                TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
                Value:          []byte(message),
            }, nil)
        }

        elapsed := time.Since(start)
        expected := time.Duration(i+1) * time.Second

        if expected > elapsed {
            sleepDuration := expected - elapsed
            fmt.Printf("Sleeping for %d ns to throttle the send speed...\n", sleepDuration)
            time.Sleep(sleepDuration)
        }

        totalSent := (i+1)*messagesPerBatch
        perSecond := int(float64(messagesDelivered) / elapsed.Seconds())
        fmt.Printf("Sent out a total of %d messages within %.2f seconds (delivered %d%%, average %d messages/sec)\n", totalSent, elapsed.Seconds(), 100 * messagesDelivered / totalSent, perSecond)
    }

    fmt.Printf("Waiting for message delivery (flush)...\n")
    beforeFlush := time.Now()
    outstanding := p.Flush(15_000)
    p.Close()

    fmt.Printf("Collected delivery report of %d messages after %.2f seconds (delivered %d%%, %d outstanding).\n", messagesDelivered, time.Since(beforeFlush).Seconds(), 100 * messagesDelivered / (messagesPerBatch * nBatches), outstanding)
}
