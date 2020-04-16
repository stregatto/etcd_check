#!/bin/bash

CERTFOLDER=$(dirname "$0")
CADURATION=3600

if [ $# -eq 0 ]
then
    echo "Please provide all nodes"
    exit 1
fi

openssl genrsa -out ${CERTFOLDER}/ca-key.pem 4096 > /dev/null
openssl req -x509 -new -nodes -key ${CERTFOLDER}/ca-key.pem -days ${CADURATION} -out ${CERTFOLDER}/ca.pem -subj "/CN=etcd-ca"

for NODE in $@
do
 openssl genrsa -out ${CERTFOLDER}/${NODE}-key.pem 2048
 openssl req -new -key ${CERTFOLDER}/${NODE}-key.pem -out ${CERTFOLDER}/${NODE}.csr -subj "/CN=${NODE}" -config ${CERTFOLDER}/cluster.conf
 openssl x509 -req -in ${CERTFOLDER}/${NODE}.csr -CA ${CERTFOLDER}/ca.pem -CAkey ${CERTFOLDER}/ca-key.pem -CAcreateserial -out ${CERTFOLDER}/${NODE}.pem -days 36000 -sha256 -extensions ssl_client -extfile ${CERTFOLDER}/cluster.conf > /dev/null 2>&1
done
