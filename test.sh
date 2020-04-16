#!/bin/bash
export CERTIFICATES=$(pwd)/etcd_tls
curl --cacert ${CERTIFICATES}/ca.pem --cert ${CERTIFICATES}/server.pem  --key ${CERTIFICATES}/server-key.pem   -L https://etcd2:2381/version

