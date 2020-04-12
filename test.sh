#!/bin/bash
export CERTIFICATES=$(pwd)/etcd_tls
curl --cacert ${CERTIFICATES}/ca.pem --cert ${CERTIFICATES}/server.pem  --key ${CERTIFICATES}/server-key.pem   -L https://etcd1:2379/version

