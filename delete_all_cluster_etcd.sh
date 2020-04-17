#!/bin/bash

#remove old run
for etcd_id in  etcd1 etcd2 etcd3
do
    docker rm -f ${etcd_id}
done

# remove old etcd-volumes
for volume in data_etcd1 data_etcd2 data_etcd3
do
    docker volume rm ${volume}
done
