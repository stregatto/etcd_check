package main

import (
	"flag"
	"fmt"
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
	// TODO:cert
	// change names accordingly to etcd library
	cacert := flag.String("cacert", "etcd_tls/ca.pem", "Certification authority file")
	cert := flag.String("cert", "etcd_tls/server.pem", "Server or Client certificate file")
	key := flag.String("key", "etcd_tls/server-key.pem", "Server or Client certificate key file")
	endPoints := flag.String("endpoints", "etcd1:2379", "Comma separated list of ETCD endpoints without protocol")
	// protocol := flag.String("proto", "https://", "Add transport protocol for http connection")
	flag.Parse()
	tlsConfig := etcd.SecureCfg{
		Cert:   *cert,
		Key:    *key,
		CaCert: *cacert,
	}

	cli := etcd.GrpcClient(tlsConfig, strings.Split(*endPoints, ","))

	// I can retrive some data
	fmt.Println(etcd.GetRaftIndexPerMembers(cli))

	fmt.Println(etcd.GetEndPointsFromInitiatedClient(cli))

	fmt.Println(etcd.GetClusterEndpoints(cli))

	//var eh epHealth
	//if err != nil {
	//	eh = epHealth{Ep: endPoints[0], Health: false, Error: err.Error()}
	//}
	//st := time.Now()
	//_, err = cli.Get(context.Background(), "health")
	//resp, err := cli.Status(context.Background(), endPoints[0])
	//
	//eh = epHealth{Ep: endPoints[0], Health: false, Took: time.Since(st).String()}
	//if err == nil {
	//	eh.Health = true
	//}

	//fmt.Println(eh)
	//fmt.Println(resp)

	//
	//Not used: is an example for v2 API
	//transport := etcd.Transport(*cacert, *cert, *key)
	//ep := *protocol + endPoints[0]
	//fmt.Println(ep)
	//etcd.Client(transport, ep)
	//
}
