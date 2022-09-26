#!/bin/bash

if [ ! -n "${1}" ];then
    echo "$0 <host>"
    exit 0
fi

HOST=$1

# https://cloud.tencent.com/developer/article/1670205
docker run -d -p 6379:6379 --name some-redis redis