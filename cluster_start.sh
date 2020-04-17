#!/bin/bash

# set common vars
LOCALPATH=$(dirname $0)
export ETCD_REGISTRY=quay.io/coreos/etcd
export ETCD_CERTIFICATES=$(pwd)/etcd_tls


# Cluster config
export DATA_DIR=/var/lib/etcd 

# For each machine
export ETCD_VERSION=latest
export TOKEN=my-etcd-token
export CLUSTER_STATE=new
NAME[1]="etcd1,172.19.0.10,2379,2380"
NAME[2]="etcd2,172.19.0.11,2381,2382"
NAME[3]="etcd3,172.19.0.12,2383,2384"


if [ -z "$1" ]
then
    echo "please give me your local IP"
    exit 1
fi

MyIP=$1

# Prepare openssl conf template
HOST_1=$(echo ${NAME[1]}|cut -d',' -f 2)
HOST_2=$(echo ${NAME[2]}|cut -d',' -f 2)
HOST_3=$(echo ${NAME[3]}|cut -d',' -f 2)
cat ${LOCALPATH}/etcd_tls/cluster.tpl |\
    sed -e "s/\${HOST_1}/${HOST_1}/" |\
    sed -e "s/\${HOST_2}/${HOST_2}/" |\
    sed -e "s/\${HOST_3}/${HOST_3}/" |\
    sed -e "s/\${MyIP}/${MyIP}/" > ${LOCALPATH}/etcd_tls/cluster.conf

# Prepare certificates
NAME_1=$(echo ${NAME[1]}|cut -d',' -f 1)
NAME_2=$(echo ${NAME[2]}|cut -d',' -f 1)
NAME_3=$(echo ${NAME[3]}|cut -d',' -f 1)

${LOCALPATH}/etcd_tls/generate_etcd_membre_cert.sh ${NAME_1} ${NAME_2} ${NAME_3}

# Delete all old clusters
./delete_all_cluster_etcd.sh 

# Configure the cluster initial var
export CLUSTER=${NAME_1}=https://${HOST_1}:2380,${NAME_2}=https://${HOST_2}:2380,${NAME_3}=https://${HOST_3}:2380


#spawn the servers
for H in 1 2 3
do
    ./node_start.sh ${NAME[${H}]} &
done
