# etcd_check
This is a simple _ETCD_ check, it's intended to verify if all endpoints are ok, and it's useful to better understand ETCD api :)

# ETCD server setup

It's useful to have an _ETCD_ server to test the software. In this repo you'll find some useful (and very "stupid") scripts to run a dummy _ETCD_ cluster with _TLS_ enabled on docker.

To use that scripts you need to source some envar like below

```
export REGISTRY=quay.io/coreos/etcd
export CERTIFICATES=The_tath_contains_the_certificates/etcd_tls
export ETCD_CONFIG_DIR=The_Path_Contains_the_etcd.env_file
```

 * `etcd.env`: contains the environment variables to run the server
 * `jumpin.sh`: I'm too lazy to write commands, it jumps in the _ETCD_ container, it's very dumb.
 * `start_etcd.sh`: It's no more than a list of commands to initiate an _ETCD_ docker container, *Note*: it deletes the previous _ETCD_ instance.
 * `test.sh` : It performs a cURL to your client.
 
Please *DO NOT USE* this scripts in production.