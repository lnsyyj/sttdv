# sttdv
storage test tool data visualization

# Install MariaDB （安装MariaDB）

```
sudo yum install -y mariadb-server && sudo systemctl start mariadb && sudo systemctl enable mariadb && sudo systemctl status mariadb
```

# onfiguring MariaDB（配置）

Authorize root user to connect remotely（授权root用户远程连接）

Replace `<YouPassword>`（替换`<YouPassword>`）

```
# Any host connects to the MariaDB server via root/<YouPassword> （任何主机通过root/<YouPassword>连接MariaDB服务器）

UPDATE mysql.user SET Password=PASSWORD('<YouPassword>') where USER='root';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY '<YouPassword>' WITH GRANT OPTION;
FLUSH PRIVILEGES;

# Example（例如）
[root@ceph-dev ~]# mysql -uroot -p
Enter password:

MariaDB [(none)]> UPDATE mysql.user SET Password=PASSWORD('1234567890') where USER='root';
Query OK, 5 rows affected (0.00 sec)
Rows matched: 5  Changed: 5  Warnings: 0

MariaDB [(none)]> GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY '1234567890' WITH GRANT OPTION;
Query OK, 0 rows affected (0.00 sec)

MariaDB [(none)]> FLUSH PRIVILEGES;
Query OK, 0 rows affected (0.00 sec)

# Create a database, create a table（创建数据库，创建表）
# Create table SQL can be found in the project Tools directory（创建表SQL在项目Tools目录中可以找到）
MariaDB [(none)]> CREATE DATABASE IF NOT EXISTS ceph DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
Query OK, 1 row affected (0.00 sec)

MariaDB [(none)]> use ceph;
Database changed

MariaDB [ceph]> SET FOREIGN_KEY_CHECKS=0;
Query OK, 0 rows affected (0.00 sec)

MariaDB [ceph]> CREATE TABLE `vdbench_filesystem` (
    ->   `Id` bigint(20) NOT NULL AUTO_INCREMENT,
    ->   `DateTime` datetime DEFAULT NULL,
    ->   `OutputInterval` int(255) DEFAULT NULL,
    ->   `ReqstdOpsRate` double(255,3) DEFAULT NULL,
    ->   `ReqstdOpsResp` double(255,3) DEFAULT NULL,
    ->   `CpuTotal` double(255,3) DEFAULT NULL,
    ->   `CpuSys` double(255,3) DEFAULT NULL,
    ->   `ReadPct` double(255,3) DEFAULT NULL,
    ->   `ReadRate` double(255,3) DEFAULT NULL,
    ->   `ReadResp` double(255,3) DEFAULT NULL,
    ->   `WriteRate` double(255,3) DEFAULT NULL,
    ->   `WriteResp` double(255,3) DEFAULT NULL,
    ->   `MbSecRead` double(255,3) DEFAULT NULL,
    ->   `MbSecWrite` double(255,3) DEFAULT NULL,
    ->   `MbSecTotal` double(255,3) DEFAULT NULL,
    ->   `XferSize` double(255,3) DEFAULT NULL,
    ->   `MkdirRate` double(255,3) DEFAULT NULL,
    ->   `MkdirResp` double(255,3) DEFAULT NULL,
    ->   `RmdirRate` double(255,3) DEFAULT NULL,
    ->   `RmdirResp` double(255,3) DEFAULT NULL,
    ->   `CreateRate` double(255,3) DEFAULT NULL,
    ->   `CreateResp` double(255,3) DEFAULT NULL,
    ->   `OpenRate` double(255,3) DEFAULT NULL,
    ->   `OpenResp` double(255,3) DEFAULT NULL,
    ->   `CloseRate` double(255,3) DEFAULT NULL,
    ->   `CloseResp` double(255,3) DEFAULT NULL,
    ->   `DeleteRate` double(255,3) DEFAULT NULL,
    ->   `DeleteResp` double(255,3) DEFAULT NULL,
    ->   `OperationTableDate` datetime DEFAULT NULL,
    ->   `TestCase` varchar(255) DEFAULT NULL,
    ->   `ClientNumber` varchar(255) DEFAULT NULL,
    ->   PRIMARY KEY (`Id`)
    -> ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
Query OK, 0 rows affected (0.00 sec)

MariaDB [ceph]> exit
Bye

```

# Install Grafana（安装Grafana）
```
wget https://dl.grafana.com/oss/release/grafana-6.2.2-1.x86_64.rpm && sudo yum localinstall grafana-6.2.2-1.x86_64.rpm

sudo systemctl start grafana-server && sudo systemctl enable grafana-server.service && sudo systemctl status grafana-server
```

# 
```cassandraql
./sttdv -ClientNumber ceph-1 -TestCase FileSystem -logPath ./863.log -mariaDBHostIP 10.121.9.23 -mariaDBDatabase cephtest -mariaDBTableName vdbench_filesystem -mariaDBUserName root -mariaDBUserPassword 1234567890 -outputinterval 1 -visualizationType vdbench
```