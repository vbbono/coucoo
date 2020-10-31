#!/bin/bash
docker network create --gateway 10.1.0.1 --subnet 10.1.0.0/16 coucoo_network
docker run --rm --network coucoo_network --ip 10.1.0.5 --name coucoo_mysql_1 -v ~/datadir:/var/lib/mysql: -e MYSQL_ROOT_PASSWORD=password -d -p 3306:3306 mysql:5.7
