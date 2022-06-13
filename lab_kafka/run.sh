#!/bin/bash

# 创建网络
docker network create app-tier --driver bridge

# 启动zookeeper
docker run -d --name zookeeper-server \
    --network app-tier \
    -e ALLOW_ANONYMOUS_LOGIN=yes \
    bitnami/zookeeper:latest

# 启动kafka
docker run -d --name kafka-server \
    --network app-tier \
    -e ALLOW_PLAINTEXT_LISTENER=yes \
    -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper-server:2181 \
    bitnami/kafka:latest

docker run -it --rm \
    --network app-tier \
    -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper-server:2181 \
    bitnami/kafka:latest kafka-topics.sh --list  --bootstrap-server kafka-server:9092