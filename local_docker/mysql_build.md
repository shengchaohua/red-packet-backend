# Build a single-node mysql with Docker
Run this command:
```shell
docker run -p 6606:3306 -e MYSQL_ROOT_PASSWORD='123456' \
  -v $(pwd)/scripts:/scripts \
  --name=red_packet_backend_mysql \
  -d mysql:5.7-debian \
  --character-set-server=utf8mb4
```

# Build a master-slave mysql cluster with Docker
> https://hub.docker.com/_/mysql

## 1. Run two mysql containers as master and slave nodes
Note: the following cmd should be run in `Project_ROOT_PATH` because we need to mount `script` directory.

Run a mysql container as the master node in port=6606:
```shell
docker run -p 6606:3306 -e MYSQL_ROOT_PASSWORD='123456' \
  -v $(pwd)/scripts:/scripts \
  --name=red_packet_backend_mysql_master \
  -d mysql:5.7-debian \
  --character-set-server=utf8mb4
```

Run a mysql container as the slave node  in port=6607:
```shell
docker run -p 6607:3306 -e MYSQL_ROOT_PASSWORD='123456' \
  -v $(pwd)/scripts:/scripts \
  --name=red_packet_backend_mysql_slave \
  -d mysql:5.7-debian  \
  --character-set-server=utf8mb4
```

## 2. Set up mysql
### 2.1. Set up mysql master node

Enter mysql master container:
```shell
docker exec -it red_packet_backend_mysql_master /bin/bash
```

Go to the mysql config path:
```shell
cd /etc/mysql/mysql.conf.d/
```

Use Vim to append the following text in the end of the `mysqld.cnf` file:
```text
server-id=100
log-bin=mysql-bin
```
Note:
- server-id is to indicate the current mysql server
- log-bin is to enable bin log 

If Vim is not available, you need to update apt and install Vim
```shell
apt-get update
apt-get install vim
```

Or use the following cmd directly:
```shell
echo "server-id=100" >> mysqld.cnf
echo "log-bin=mysql-bin" >> mysqld.cnf
````

### 3.2 Set up mysql slave node
Enter mysql slave container:
```shell
docker exec -it red_packet_backend_mysql_slave /bin/bash
```

Modify the same mysql config file by appending these text:
```text
server-id=101
relay-log=mysql-relay-log
```

Note:
- server-id is to indicate the current mysql server
- relay-log is to enable relay log

Or use the following cmd directly:
```shell
echo "server-id=101" >> mysqld.cnf 
echo "relay-log=mysql-relay-log" >> mysqld.cnf
````

### 3.3 Set up master-slave syncing
#1. Restart both containers
```shell
docker restart red_packet_backend_mysql_master red_packet_backend_mysql_slave
```

#2. Get the IP address of both containers:

Open a new terminal and run these cmd:
```shell
docker inspect --format='{{.NetworkSettings.IPAddress}}' red_packet_backend_mysql_master
docker inspect --format='{{.NetworkSettings.IPAddress}}' red_packet_backend_mysql_slave
```

Assume we get two IP address:
- master node: 172.17.0.2
- slave node: 172.17.0.3

#3. Set up mysql master node again

Enter mysql master shell
```shell
docker exec -it red_packet_backend_mysql_master mysql -uroot -p123456
```

Run these cmd in mysql shell to create a new user:
```mysql
CREATE USER 'slave'@'%' IDENTIFIED BY '123456';
GRANT REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'slave'@'%';
```

Check master status in mysql shell:
```mysql
show master status;
```

Assume we get:
```text
+------------------+----------+--------------+------------------+-------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
+------------------+----------+--------------+------------------+-------------------+
| mysql-bin.000001 |      154 |              |                  |                   |
+------------------+----------+--------------+------------------+-------------------+
1 row in set (0.00 sec)
```

The first two fields will be used later:
- File: mysql-bin.000001
- Position: 154

#5. Setup mysql slave node again

Enter mysql slave shell:
```shell
docker exec -it red_packet_backend_mysql_slave mysql -uroot -p123456
```

Run this cmd with replacing 
- master_log_file -> File obtained in the last step
- master_log_pos -> Position obtained in the last step
- master_host -> the IP address of master node

```mysql
CHANGE MASTER TO master_host = '172.17.0.2',
master_user = 'slave',
master_password = '123456',
master_port = 3306,
master_log_file = 'mysql-bin.000001',
master_log_pos = 154,
master_connect_retry = 30;
```

## 4. Test
#1. Restart both containers
```shell
docker restart red_packet_backend_mysql_master red_packet_backend_mysql_slave
```

#2. Create a new database in mysql master node
```mysql
CREATE DATABASE test_db;
```

#3. Check it in mysql slave node
```mysql
SHOW databases;
```

If you find test_db database in slave node, it means you have just succeeded to build a master-slave mysql cluster.