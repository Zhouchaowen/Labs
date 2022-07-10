基础环境

```bash
yum install python3 python3-devel gcc -y
```

PSQL

```bash
sudo yum install -y https://download.postgresql.org/pub/repos/yum/reporpms/EL-6-x86_64/pgdg-redhat-repo-latest.noarch.rpm
sudo yum install -y postgresql12-server

# sudo service postgresql-12 initdb
sudo /usr/pgsql-12/bin/postgresql-12-setup initdb
sudo chkconfig postgresql-12 on
sudo service postgresql-12 start
```

```bash
# vim pg_hba.conf 
host    all             all             0.0.0.0/0               md5
host    replication     all             0.0.0.0/0               md5
```

```bash
# vim postgresql.conf 
listen_addresses = '*'
hot_standby = on
log_connections = on
log_disconnections = on
log_statement = 'all'
```

```bash
# 创建用户设置密码
su - postgres
create role repl with replaction login
\password repl
```

```
systemctl enable postgresql-12
```



```
pg_basebackup -h 10.2.0.104 -D data -U repl -P -v -R -X stream -S repl2 -C

pg_basebackup -h 10.2.0.104 -D data -U repl -P -v -R -X stream -S repl3 -C
```



ETCD

```bash
wget -c https://github.com/etcd-io/etcd/releases/download/v3.4.7/etcd-v3.4.7-linux-amd64.tar.gz

tar xvf etcd-v3.4.7-linux-amd64.tar.gz
mv etcd-v3.4.7-linux-amd64 etcd-v3.4.7
mv etcd-v3.4.7 /opt
```

```bash
#!/bin/bash
etcd --name etcd_ydtf01 \
  --initial-advertise-peer-urls http://10.2.0.104:2380 \
  --listen-peer-urls http://10.2.0.104:2380 \
  --listen-client-urls http://10.2.0.104:2379,http://127.0.0.1:2379 \
  --advertise-client-urls http://10.2.0.104:2379 \
  --initial-cluster-token etcd-cluster-ydtf \
  --initial-cluster etcd_ydtf01=http://10.2.0.104:2380,etcd_ydtf02=http://10.2.0.105:2380,etcd_ydtf03=http://10.2.0.106:2380 \
  --initial-cluster-state new \
  --enable-v2
```

```bash
#!/bin/bash
etcd --name etcd_ydtf02 \
  --initial-advertise-peer-urls http://10.2.0.105:2380 \
  --listen-peer-urls http://10.2.0.105:2380 \
  --listen-client-urls http://10.2.0.105:2379,http://127.0.0.1:2379 \
  --advertise-client-urls http://10.2.0.105:2379 \
  --initial-cluster-token etcd-cluster-ydtf \
  --initial-cluster etcd_ydtf01=http://10.2.0.104:2380,etcd_ydtf02=http://10.2.0.105:2380,etcd_ydtf03=http://10.2.0.106:2380 \
  --initial-cluster-state new \
  --enable-v2
```

```bash
#!/bin/bash
etcd --name etcd_ydtf03 \
  --initial-advertise-peer-urls http://10.2.0.106:2380 \
  --listen-peer-urls http://10.2.0.106:2380 \
  --listen-client-urls http://10.2.0.106:2379,http://127.0.0.1:2379 \
  --advertise-client-urls http://10.2.0.106:2379 \
  --initial-cluster-token etcd-cluster-ydtf \
  --initial-cluster etcd_ydtf01=http://10.2.0.104:2380,etcd_ydtf02=http://10.2.0.105:2380,etcd_ydtf03=http://10.2.0.106:2380 \
  --initial-cluster-state new \
  --enable-v2
```



Patroni

```
yum install postgresql-devel -y
```

```
pip3 install psycopg2

pip3 install patroni[etcd] -i https://mirrors.aliyun.com/pypi/simple/
pip3 install psycopg2-binary -i https://mirrors.aliyun.com/pypi/simple/
pip3 install patroni -i https://mirrors.aliyun.com/pypi/simple/
```

```yaml
scope: pg_ydtf
namespace: /service/
name: pg_ydtf01

restapi:
  listen: 10.2.0.104:8008
  connect_address: 10.2.0.104:8008

etcd:
  host: 10.2.0.104:2379

bootstrap:
  dcs:
    ttl: 30
    loop_wait: 10
    retry_timeout: 10
    maximum_lag_on_failover: 1048576
    master_start_timeout: 300
    synchronous_mode: false
    postgresql:
      use_pg_rewind: true
      use_slots: true
      parameters:
      wal_level: locical
      hot_standby: "on"
      wal_keep_segments: 128
      max_wal_senders: 10
      max_replication_slots: 10
      wal_log_hints: "on"
      archive_mode: "on"
      hot_standby: on
      archive_timeout: 1800s

postgresql:
  listen: 0.0.0.0:5432
  connect_address: 10.2.0.104:5432
  data_dir: /var/lib/pgsql/12/data
  bin_dir: /usr/pgsql-12/bin/
  config_dir: /var/lib/pgsql/12/data
  authentication:
    replication:
      username: repl
      password: zdns
    superuser:
      username: postgres
      password: postgres
    rewind:  # Has no effect on postgres 10 and lower
      username: repl
      password: zdns

tags:
    nofailover: false
    noloadbalance: false
    clonefrom: false
    nosync: false
```

```yaml
scope: pg_ydtf
namespace: /service/
name: pg_ydtf02

restapi:
  listen: 10.2.0.105:8008
  connect_address: 10.2.0.105:8008

etcd:
  host: 10.2.0.105:2379

bootstrap:
  dcs:
    ttl: 30
    loop_wait: 10
    retry_timeout: 10
    maximum_lag_on_failover: 1048576
    master_start_timeout: 300
    synchronous_mode: false
    postgresql:
      use_pg_rewind: true
      use_slots: true
      parameters:
      wal_level: locical
      hot_standby: "on"
      wal_keep_segments: 128
      max_wal_senders: 10
      max_replication_slots: 10
      wal_log_hints: "on"
      archive_mode: "on"
      hot_standby: on
      archive_timeout: 1800s

postgresql:
  listen: 0.0.0.0:5432
  connect_address: 10.2.0.105:5432
  data_dir: /var/lib/pgsql/12/data
  bin_dir: /usr/pgsql-12/bin/
  config_dir: /var/lib/pgsql/12/data
  authentication:
    replication:
      username: repl
      password: zdns
    superuser:
      username: postgres
      password: postgres
    rewind:  # Has no effect on postgres 10 and lower
      username: repl
      password: zdns

tags:
    nofailover: false
    noloadbalance: false
    clonefrom: false
    nosync: false
```

```yaml
scope: pg_ydtf
namespace: /service/
name: pg_ydtf03

restapi:
  listen: 10.2.0.106:8008
  connect_address: 10.2.0.106:8008

etcd:
  host: 10.2.0.106:2379

bootstrap:
  dcs:
    ttl: 30
    loop_wait: 10
    retry_timeout: 10
    maximum_lag_on_failover: 1048576
    master_start_timeout: 300
    synchronous_mode: false
    postgresql:
      use_pg_rewind: true
      use_slots: true
      parameters:
      wal_level: locical
      hot_standby: "on"
      wal_keep_segments: 128
      max_wal_senders: 10
      max_replication_slots: 10
      wal_log_hints: "on"
      archive_mode: "on"
      hot_standby: on
      archive_timeout: 1800s

postgresql:
  listen: 0.0.0.0:5432
  connect_address: 10.2.0.106:5432
  data_dir: /var/lib/pgsql/12/data
  bin_dir: /usr/pgsql-12/bin/
  config_dir: /var/lib/pgsql/12/data
  authentication:
    replication:
      username: repl
      password: zdns
    superuser:
      username: postgres
      password: postgres
    rewind:  # Has no effect on postgres 10 and lower
      username: repl
      password: zdns

tags:
    nofailover: false
    noloadbalance: false
    clonefrom: false
    nosync: false
```



# 参考

https://postgres.fun/20200529182600.html

https://zhuanlan.zhihu.com/p/260958352

https://www.postgresql.org/download/linux/redhat/

https://www.modb.pro/topic/152353