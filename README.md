# Confluent Kafka Go: Simple Performance Measurement

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

On my machine, I send out 30 million messages within ~30 seconds, but it takes
me around 240 seconds to receive all 30 million messages.

For a complete log on my machine, see [results.log](results.log).
