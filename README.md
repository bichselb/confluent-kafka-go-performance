# Confluent Kafka Go: Measuring Serialization Performance

This code measures the performance of the [confluent kafka
go](https://docs.confluent.io/platform/current/clients/confluent-kafka-go/index.html)
library when serializing using protobuf and the schema registry, as showcased
[here](https://github.com/confluentinc/confluent-kafka-go/blob/master/examples/protobuf_producer_example/protobuf_producer_example.go).

## Run

Install docker and run the experiment with:

```bash
./run.sh 2>&1 | tee results.log
```

## Results

On my machine, I get:

| Test                     | Time to serialize 1 million messages |
| ------------------------ | ------------------------------------ |
| Protobuf+schema registry | 1.603282821s                         |
| Pure protobuf (baseline) | 99.182533ms                          |

For a complete log on my machine, see [results.log](results.log).
