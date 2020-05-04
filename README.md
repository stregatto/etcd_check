# etcd_check
This is a simple _ETCD_ check, it's intended to verify if all endpoints are ok, and it's useful to better understand ETCD api :)

## Long story short

To run this software you have to enumerate all endpoints and the certificate it will use to access them

`etcd_check -endpoints etcd1:2379,etcd2:2381,etcd3:2383 -cert ./etcd_tls/etcd1.pem -key ./etcd_tls/etcd1-key.pem -n -u`

## Usage

```
etcd_check:
  -cacert string
        Certification authority file (default "etcd_tls/ca.pem")
  -cert string
        Server or Client certificate file (default "etcd_tls/server.pem")
  -endpoints string
        Comma separated list of ETCD endpoints without protocol (default "etcd1:2379")
  -key string
        Server or Client certificate key file (default "etcd_tls/server-key.pem")
  -maxFailingMember int
        The max number of ETCD servers can fail in your cluster
  -maxRaftDrift int
        The max drift the raft index can support (default 1)
  -n    Print cluster status in NAGIOS ready format
  -u    Print only unreachable nodes in NAGIOS ready format, it excludes raft check
```

## Returns

It returns the status of the cluster depending on flags.

### -n

`-n` returns the status of cluster in _NAGIOS_ format. The RAFT index of all cluster's members are checked, all RAFT indexes must be in the +/- `maxRaftDrift` interval.

### -u
`-u` returns the status of cluster in _NAGIOS_ format. The RAFT is not checked, it fails if the number of failing (not reached) members is more than `maxFailingMember`.


# Compiling issues

Because of https://github.com/etcd-io/etcd/issues/11154 you need some tricks to compile your go.mod.

In your go.etcd.io checkout v3.4 perform the below command to retrieve the MVS (Minimal Version Selection) and with that you can create your _requirement_
```
TZ=UTC git --no-pager show \
>   --quiet \
>   --abbrev=12 \
>   --date='format-local:%Y%m%d%H%M%S' \
>   --format="%cd-%h"
```

In go.mod set this:
```
require go.etcd.io/etcd v0.0.0-[THE_OUTPUT_OF_ABOVE_COMMAND]
```

# Just for experiment

## ETCD server setup

It's useful to have an _ETCD_ cluster to test the software. In this repo you'll find some useful (and very "stupid") scripts to run a dummy _ETCD_ cluster with _TLS_ on docker.

 * `cluster_start.sh`: It's no more than a list of commands to initiate an _ETCD_ cluster composed by 3 ETCD nodes, *Note*: it deletes the previous _ETCD_ instance. Use it to spawn your cluster
 * `node_start.sh`: It's called by `cluster_start.sh` to initiate the _ETCD_ docker containers
 * `delete_all_cluster_etcd.sh`: It's called by `cluster_start.sh` to delete all previously spawned _ETCD_ instance
 * `test.sh` : It performs a cURL to your first etcd server via docker exported port and perform etcdctl cluster status to your first _etcd1_ .
 * `jumpin.sh`: I'm too lazy to write commands, it jumps in the first _ETCD_ container, it's very dumb.
 
Please *DO NOT USE* the .sh scripts in production.
