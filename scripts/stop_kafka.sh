#!/bin/bash

#kill zookeeper
echo "===> stop zookeeper service"
ps ax | grep 'zookeeper' | awk '{print $1}' | xargs kill

echo "===> stop kafka service"
ps ax | grep 'kafka-server' | awk '{print $1}' | xargs kill
