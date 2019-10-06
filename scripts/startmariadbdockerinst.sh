#!/bin/bash
docker run --name dbmaria -e MYSQL_ROOT_PASSWORD=maria -p 3306:3306 -d mariadb:latest 