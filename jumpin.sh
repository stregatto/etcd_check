#!/bin/bash

docker exec -t -i `docker ps |grep etcd|awk '{ print $1 }'` /bin/sh
