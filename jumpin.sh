#!/bin/bash

docker exec -t -i `docker ps |grep etcd|head -n1|awk '{ print $1 }'` /bin/sh
