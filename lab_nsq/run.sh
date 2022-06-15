#!/bin/bash

if [ ! -n "${1}" ];then
    echo "$0 <host>"
    exit 0
fi

HOST=$1

# 启动 nsqlookup
docker run -d --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq /nsqlookupd

docker run -d --name nsqd -p 4150:4150 -p 4151:4151 \
    nsqio/nsq /nsqd \
    --broadcast-address=${HOST} \
    --lookupd-tcp-address=${HOST}:4160

