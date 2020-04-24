package etcd

import (
	"go.etcd.io/etcd/etcdserver/etcdserverpb"
)

type commandResChan struct {
	server string
	res    interface{}
	err    error
}

type ClusterEndpoints struct {
	Server string
	Ep     []*etcdserverpb.Member
	Err    error
}

type ClientEndpoints struct {
	Server string
	Ep     []string
	Err    error
}

type RaftIndexPerMember struct {
	Server    string
	MemberId  uint64
	RaftIndex uint64
	Err       error
}
