#!/bin/bash
export CERTIFICATES=$(pwd)/etcd_tls

echo "curl to docker exported port" 
curl --cacert ${CERTIFICATES}/ca.pem --cert ${CERTIFICATES}/etcd1.pem  --key ${CERTIFICATES}/etcd1-key.pem   -L https://etcd1:2381/version

echo "testing from first docker container"

docker exec -t -i etcd1 /bin/sh -c 'ETCDCTL_API=3 /usr/local/bin/etcdctl --endpoints=etcd1:2379 --cacert=/etc/ssl/etcd_tls/ca.pem --cert=/etc/ssl/etcd_tls/etcd1.pem --key=/etc/ssl/etcd_tls/etcd1-key.pem -w table endpoint --cluster status'


