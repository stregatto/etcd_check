#!/bin/bash
export CERTIFICATES=$(pwd)/etcd_tls
curl --cacert ${CERTIFICATES}/ca.pem --cert ${CERTIFICATES}/etcd1.pem  --key ${CERTIFICATES}/etcd1-key.pem   -L https://etcd1:2381/version

