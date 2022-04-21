#!/bin/bash

PORT=9092
KAFKA_TOPIC=product

# start zookeeper service
echo "===> start zookeeper service"
./kafka/bin/zookeeper-server-start.sh ./configs/zookeeper.properties 2>&1>build/kafka/logs/zookeeper.log &

sleep 10

echo "===> start kafka service"
./kafka/bin/kafka-server-start.sh ./configs/server.properties 2>&1>build/kafka/logs/server.log &

sleep 5

echo "===> create producer topic"
./kafka/bin/kafka-topics.sh --create --topic $KAFKA_TOPIC --bootstrap-server localhost:$PORT -partitions 4