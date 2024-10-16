#!/bin/bash

# enable bash strict mode
# http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -euo pipefail

echo "Compose down..."
docker compose down --remove-orphans --volumes

echo "Compose up..."
docker compose up --build -d

echo "Wait to fully start..."
sleep 10

echo "Creating kafka topic..."
docker compose exec broker /usr/bin/kafka-topics --create \
  --topic yourtopic \
  --bootstrap-server broker:29092 \
  --replication-factor 1 \
  --partitions 1

echo "Running go..."
docker compose exec go sh -c "cd /host && protoc --go_out=. --go-grpc_out=. mymessage.proto && cd /host/go_example && go mod download && go build && ./go_example"
