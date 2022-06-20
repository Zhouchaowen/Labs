#!/bin/bash

if [ ! -n "${1}" ];then
    echo "$0 <host>"
    exit 0
fi

HOST=$1

docker run -d --name Etcd-server \
    --network host \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://${HOST}:2379 \
    bitnami/etcd:latest