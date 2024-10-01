# Confluent Kafka Go: Simple Performance Measurement

## Run

Install docker and run the experiment with `./run.sh`.

## Results

On my machine, I send out 30 million messages within ~30 seconds, but it takes
me around 240 seconds to receive all 30 million messages.

For a complete log on my machine, see [results.log](results.log).
