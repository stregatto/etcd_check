#!/bin/bash

#remove old run 
docker rm `docker ps -a |grep etcd|awk '{print $1}'`

#start new run
export DATA_DIR="etcd-data"
export REGISTRY=quay.io/coreos/etcd
export CERTIFICATES=$(pwd)/etcd_tls
export ETCD_CONFIG_DIR=$(pwd)
docker run   -p 2379:2379   -p 2380:2380   --restart=on-failure:5   --env-file=${ETCD_CONFIG_DIR}/etcd.env   -v ${CERTIFICATES}:/etc/ssl/etcd_tls:ro   -v ${DATA_DIR}:/etcd-data   --name etcd ${REGISTRY}:latest   /usr/local/bin/etcd
