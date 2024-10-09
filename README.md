# Confluent Kafka Go: Simple Performance Measurement

This repository performs a simple performance measurement of the [confluent
kafka
go](https://docs.confluent.io/platform/current/clients/confluent-kafka-go/index.html)
library.

## Run

Install docker and run the experiment with:

```bash
./run.sh 2>&1 | tee results.log
```

## Code

The core pieces of this repo are:

- The [producer](go-scripts/producer.go)
- The [consumer](go-scripts/consumer.go)
- The [docker compose file](docker-compose.yml) specifying kafka

## Results

On my machine, I get:

| Service  | Throughput           |
| -------- | -------------------- |
| Producer | 1011236 messages/sec |
| Consumer |  803654 messages/sec |

For a complete log on my machine, see [results.log](results.log).
