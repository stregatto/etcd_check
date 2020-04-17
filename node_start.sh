#!/bin/bash
#
#
# I need vars as environment like the below, cluster_common.sh provides these for me
# ETCD_REGISTRY=quay.io/coreos/etcd
# ETCD_CERTIFICATES=$(pwd)/etcd_tls
# ETCD_VERSION=latest
# TOKEN=my-etcd-token
# CLUSTER_STATE=new
# CLUSTER=${NAME_1}=https://${HOST_1}:2380,${NAME_2}=https://${HOST_2}:2380,${NAME_3}=https://${HOST_3}:2380
# DATA_DIR=/var/lib/etcd
#



if [ $# -eq 0 ]
then
    echo "Please I need THIS_NAME, THIS_IP, PORT2379, PORT2380 as input in comma separated value"
    exit 1
fi

IFS=',' read -r -a INPUT <<< "$1"

THIS_NAME=${INPUT[0]}
THIS_IP=${INPUT[1]}
PORT2379=${INPUT[2]}
PORT2380=${INPUT[3]}

# create volumes
docker volume create --name data_${THIS_NAME}

# For node 1
docker run \
  -p ${PORT2379}:2379 \
  -p ${PORT2380}:2380 \
  -h ${THIS_NAME} \
  --net docker_default \
  --ip ${THIS_IP} \
  --volume=data_${THIS_NAME}:/etcd-data \
  --volume=${ETCD_CERTIFICATES}:/etc/ssl/etcd_tls:ro \
  --name=${THIS_NAME} ${ETCD_REGISTRY}:${ETCD_VERSION} \
  /usr/local/bin/etcd \
  --data-dir /etcd-data \
  --name ${THIS_NAME} \
  --initial-advertise-peer-urls https://${THIS_IP}:2380 \
  --listen-peer-urls https://${THIS_IP}:2380 \
  --advertise-client-urls https://${THIS_IP}:2379 \
  --listen-client-urls https://${THIS_IP}:2379,https://127.0.0.1:2379 \
  --initial-cluster ${CLUSTER} \
  --initial-cluster-state ${CLUSTER_STATE} \
  --initial-cluster-token ${TOKEN} \
  --trusted-ca-file /etc/ssl/etcd_tls/ca.pem \
  --cert-file /etc/ssl/etcd_tls/${THIS_NAME}.pem \
  --key-file /etc/ssl/etcd_tls/${THIS_NAME}-key.pem \
  --client-cert-auth=true \
  --peer-trusted-ca-file /etc/ssl/etcd_tls/ca.pem \
  --peer-cert-file /etc/ssl/etcd_tls/${THIS_NAME}.pem \
  --peer-key-file /etc/ssl/etcd_tls/${THIS_NAME}-key.pem \
  --peer-client-cert-auth=true
