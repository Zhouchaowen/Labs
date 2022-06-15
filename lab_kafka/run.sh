#!/bin/bash

if [ ! -n "${1}" ];then
    echo "$0 <host>"
    exit 0
fi

HOST=$1

# 启动zookeeper
docker run -d --name zookeeper-server \
    --network host \
    -e ALLOW_ANONYMOUS_LOGIN=yes \
    bitnami/zookeeper:latest

# 启动kafka
docker run -d --name kafka-server \
    --network host \
    -e KAFKA_BROKER_ID=1 \
    -e KAFKA_CFG_LISTENERS=PLAINTEXT://${HOST}:9092 \
    -e KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://${HOST}:9092 \
    -e ALLOW_PLAINTEXT_LISTENER=yes \
    -e KAFKA_CFG_ZOOKEEPER_CONNECT=${HOST}:2181 \
    -e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=true \
    bitnami/kafka:latest

#docker run -it --rm \
#    --network app-tier \
#    -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper-server:2181 \
#    bitnami/kafka:latest kafka-topics.sh --create  --bootstrap-server kafka-server:9092 --replication-factor 1 --partitions 1 --topic standAlone
#
#docker run -it --rm \
#    --network app-tier \
#    -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper-server:2181 \
#    bitnami/kafka:latest kafka-topics.sh --list  --bootstrap-server kafka-server:9092