package etcd

import (
	"go.etcd.io/etcd/etcdserver/etcdserverpb"
)

type commandResChan struct {
	server string
	res    interface{}
}

type clusterEndpoints struct {
	server string
	ep     []*etcdserverpb.Member
}

type clientEndpoints struct {
	server string
	ep     []string
}

type raftIndexPerMember struct {
	server    string
	memberId  uint64
	raftIndex uint64
}
