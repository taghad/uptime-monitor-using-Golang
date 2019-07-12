# uptime-monitor-using-Golang
### Introduction
Written this task for internship.

#### Prerequisites

- MySQL
- Golang


#### Deployment
First create the monitorDB.sql.
```bash
mysql> CREATE DATABASE monitorDB;
```
create manager user.
```bash
mysql> CREATE USER 'manager'@'localhost' IDENTIFIED BY '123456';
```
granting for manager.
```bash
mysql> GRANT ALL PRIVILEGES ON monitorDB.* TO 'manager'@'localhost';
```
Then you can run the url uptime-monitor simply by go run:
```bash
 go run main.go
```
