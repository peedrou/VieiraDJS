#!/bin/bash

echo "Waiting for Kafka to start..."

# Wait for Kafka to be ready on internal Docker network
while ! nc -z kafka 9092; do
  echo "Kafka not ready, waiting..."
  sleep 1
done

echo "Kafka started!"

# Create the topic using Kafka's internal address
/kafka/bin/kafka-topics.sh --create --topic example-topic --bootstrap-server kafka:9092 --partitions 1 --replication-factor 1

echo "Topic created!"
