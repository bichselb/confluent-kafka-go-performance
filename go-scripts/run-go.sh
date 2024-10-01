#!/bin/sh

echo "Waiting for kafka to start up..."
sleep 20

echo "Running producer..."
./producer &

echo "Running consumer after slight delay..."
sleep 0.5
./consumer &

echo "Waiting for producer and consumer to finish..."
wait