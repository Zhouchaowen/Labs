#!/bin/bash

PATH="/root/patroni"
mkdir ${PATH}
cd ${PATH}

# [基础环境]
yum install vim python3 python3-devel gcc -y

# [ETCD]
wget -c https://github.com/etcd-io/etcd/releases/download/v3.4.7/etcd-v3.4.7-linux-amd64.tar.gz
tar xvf etcd-v3.4.7-linux-amd64.tar.gz
mv etcd-v3.4.7-linux-amd64 etcd-v3.4.7

# [Postgres]
sudo yum install -y https://download.postgresql.org/pub/repos/yum/reporpms/EL-6-x86_64/pgdg-redhat-repo-latest.noarch.rpm
sudo yum install -y postgresql12-server
sudo yum install -y postgresql-devel
sudo /usr/pgsql-12/bin/postgresql-12-setup initdb
sudo /usr/pgsql-12/bin/postgresql-12 start

# [Patroni]
pip3 install psycopg2
pip3 install patroni[etcd] -i https://mirrors.aliyun.com/pypi/simple/
pip3 install psycopg2-binary -i https://mirrors.aliyun.com/pypi/simple/
pip3 install patroni -i https://mirrors.aliyun.com/pypi/simple/

