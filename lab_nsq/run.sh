#!/bin/bash

# validate command
if [ ! -n "${1}" ];then
    echo "$0 <host>"
    exit 0
fi

HOST=$1
PORT=4160

# 启动 nsqlookup
docker run -it --rm --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq /nsqlookupd

docker run -it --rm --name nsqd -p 4150:4150 -p 4151:4151 \
    nsqio/nsq /nsqd \
    --broadcast-address=${HOST} \
    --lookupd-tcp-address=${HOST}:${PORT}

