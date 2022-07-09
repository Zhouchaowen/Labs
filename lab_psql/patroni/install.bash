#!/bin/bash


yum install vim python3 python3-devel gcc -y

wget -c https://github.com/etcd-io/etcd/releases/download/v3.4.7/etcd-v3.4.7-linux-amd64.tar.gz

tar xvf etcd-v3.4.7-linux-amd64.tar.gz
mv etcd-v3.4.7-linux-amd64 etcd-v3.4.7
mv etcd-v3.4.7 /opt

