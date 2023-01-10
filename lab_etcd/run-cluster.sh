#!/bin/bash

TOKEN=token-01
CLUSTER_STATE=new
NAME_1=machine-1
NAME_2=machine-2
NAME_3=machine-3
HOST_1=10.240.0.17
HOST_2=10.240.0.18
HOST_3=10.240.0.19
CLUSTER=${NAME_1}=http://${HOST_1}:2380,${NAME_2}=http://${HOST_2}:2380,${NAME_3}=http://${HOST_3}:2380

docker run -d --name Etcd-server-1 \
    --network host \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_DATA_DIR=data.etcd \
    --env ETCD_NAME=${NAME_1} \
    --env ETCD_INITIAL_ADVERTISE_PEER_URLS=http://${HOST_1}:2380 \
    --env ETCD_LISTEN_PEER_URLS=http://${HOST_1}:2380 \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://${HOST_1}:2379 \
    --env ETCD_LISTEN_CLIENT_URLS=http://${HOST_1}:2379 \
    --env ETCD_INITIAL_CLUSTER=${CLUSTER} \
    --env ETCD_INITIAL_CLUSTER_STATE=${CLUSTER_STATE} \
    --env ETCD_INITIAL_CLUSTER_TOKEN=${TOKEN} \
    bitnami/etcd:latest

docker run -d --name Etcd-server-2 \
    --network host \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_DATA_DIR=data.etcd \
    --env ETCD_NAME=${NAME_2} \
    --env ETCD_INITIAL_ADVERTISE_PEER_URLS=http://${HOST_2}:2380 \
    --env ETCD_LISTEN_PEER_URLS=http://${HOST_2}:2380 \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://${HOST_2}:2379 \
    --env ETCD_LISTEN_CLIENT_URLS=http://${HOST_2}:2379 \
    --env ETCD_INITIAL_CLUSTER=${CLUSTER} \
    --env ETCD_INITIAL_CLUSTER_STATE=${CLUSTER_STATE} \
    --env ETCD_INITIAL_CLUSTER_TOKEN=${TOKEN} \
    bitnami/etcd:latest

docker run -d --name Etcd-server-3 \
    --network host \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_DATA_DIR=data.etcd \
    --env ETCD_NAME=${NAME_3} \
    --env ETCD_INITIAL_ADVERTISE_PEER_URLS=http://${HOST_3}:2380 \
    --env ETCD_LISTEN_PEER_URLS=http://${HOST_3}:2380 \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://${HOST_3}:2379 \
    --env ETCD_LISTEN_CLIENT_URLS=http://${HOST_3}:2379 \
    --env ETCD_INITIAL_CLUSTER=${CLUSTER} \
    --env ETCD_INITIAL_CLUSTER_STATE=${CLUSTER_STATE} \
    --env ETCD_INITIAL_CLUSTER_TOKEN=${TOKEN} \
    bitnami/etcd:latest