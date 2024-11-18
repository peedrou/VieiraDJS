#!/bin/bash

echo "Waiting for Kafka to start..."
while ! nc -z localhost 9092; do
  sleep 1
done
echo "Kafka started!"

kafka-topics.sh --create --topic example-topic --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
kafka-topics.sh --create --topic another-topic --bootstrap-server localhost:9092 --partitions 2 --replication-factor 1

echo "Topics created!"
