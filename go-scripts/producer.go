package main

import (
    "fmt"
    "strconv"
    "time"

    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    start := time.Now()

    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "kafka:29092",
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
            time.Sleep(expected - elapsed)
        }

        totalSent := (i+1)*messagesPerBatch
        fmt.Printf("Sent out %d messages each within %.2f seconds (delivered %d%%)\n", totalSent, elapsed.Seconds(), 100 * messagesDelivered / totalSent)
    }

    fmt.Printf("Waiting for message deliverables...\n")
    beforeFlush := time.Now()
    p.Flush(15_000)
    fmt.Printf("Collected all message deliverables after %.2f seconds, delivering a total of %d messages.\n", time.Since(beforeFlush).Seconds(), messagesDelivered)
}
