# etcd_check
This is a simple _ETCD_ check, it's intended to verify if all endpoints are ok, and it's useful to better understand ETCD api :)

# ETCD server setup

It's useful to have an _ETCD_ cluster to test the software. In this repo you'll find some useful (and very "stupid") scripts to run a dummy _ETCD_ cluster with _TLS_ on docker.

 * `cluster_start.sh`: It's no more than a list of commands to initiate an _ETCD_ cluster composed by 3 ETCD nodes, *Note*: it deletes the previous _ETCD_ instance. Use it to spawn your cluster
 * `node_start.sh`: It's called by `cluster_start.sh` to initiate the _ETCD_ docker containers
 * `delete_all_cluster_etcd.sh`: It's called by `cluster_start.sh` to delete all previously spawned _ETCD_ instance
 * `test.sh` : It performs a cURL to your first etcd server.
 * `jumpin.sh`: I'm too lazy to write commands, it jumps in the first _ETCD_ container, it's very dumb.
 
Please *DO NOT USE* this scripts in production.