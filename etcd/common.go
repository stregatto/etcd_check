package etcd

import (
	"go.etcd.io/etcd/etcdserver/etcdserverpb"
)

type commandResChan struct {
	server string
	res    interface{}
	err    error
}

// ClusterEndpoints contains all the ETCD cluster members from the cluster point of view
type ClusterEndpoints struct {
	Server string
	Ep     []*etcdserverpb.Member
	Err    error
}

// ClientEndpoints contains all the ETCD cluster members from the client (the cli) point of view
type ClientEndpoints struct {
	Server string
	Ep     []string
	Err    error
}

// RaftIndexPerMember contains the RaftData for every member queried
type RaftIndexPerMember map[string]RaftData

// RaftData contains the RAFT index data, the member ID and the error returned by member queried
type RaftData struct {
	MemberID  uint64
	RaftIndex uint64
	Err       error
}
