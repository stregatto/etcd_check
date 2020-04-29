package main

import (
	"flag"
	"fmt"
	"github.com/stregatto/etcd_check/core"
	"github.com/stregatto/etcd_check/etcd"
	"strings"
)

type epHealth struct {
	Ep     string `json:"endpoint"`
	Health bool   `json:"health"`
	Took   string `json:"took"`
	Error  string `json:"error,omitempty"`
}

func main() {
	cacert := flag.String("cacert", "etcd_tls/ca.pem", "Certification authority file")
	cert := flag.String("cert", "etcd_tls/server.pem", "Server or Client certificate file")
	key := flag.String("key", "etcd_tls/server-key.pem", "Server or Client certificate key file")
	maxFailedMembers := flag.Int("maxFailingMember", 1, "The max number of ETCD servers can fail in your cluster")
	maxRaftDrift := flag.Int("maxRaftDrift", 1, "The max drift the raft index can support")
	endPoints := flag.String("endpoints", "etcd1:2379", "Comma separated list of ETCD endpoints without protocol")
	// protocol := flag.String("proto", "https://", "Add transport protocol for http connection")
	flag.Parse()
	tlsConfig := etcd.SecureCfg{
		Cert:   *cert,
		Key:    *key,
		CaCert: *cacert,
	}

	cli := etcd.GrpcClient(tlsConfig, strings.Split(*endPoints, ","))

	// I can retrieve some data
	fmt.Println(etcd.GetRaftIndexPerMembers(cli))
	fmt.Println(etcd.GetEndPointsFromInitiatedClient(cli))
	fmt.Println(etcd.GetClusterEndpoints(cli))
	fmt.Println(core.MembersHealthiness(etcd.GetRaftIndexPerMembers(cli), *maxFailedMembers))
	fmt.Println(core.RaftCoherence(etcd.GetRaftIndexPerMembers(cli), *maxRaftDrift))

	//Not used: is an example for v2 API

	//transport := etcd.Transport(*cacert, *cert, *key)
	//ep := *protocol + endPoints[0]
	//fmt.Println(ep)
	//etcd.Client(transport, ep)
	//
}
