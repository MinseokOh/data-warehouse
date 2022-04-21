#!/bin/bash

# Kafka Version
KAFKA_VERSION=kafka_2.13-3.1.0

echo "===> check kafka installed..."
if [ -d "./kafka" ]; then
  echo "===> already installed kafka"
else
  wget https://downloads.apache.org/kafka/3.1.0/$KAFKA_VERSION.tgz
  tar -xzf $KAFKA_VERSION.tgz

  mv $KAFKA_VERSION kafka
  rm $KAFKA_VERSION.tgz

  echo "===> installed kafka"
fi
