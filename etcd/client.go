// Package etcd contains core to access ETCD apis both v2 and v3.
// APIs v2 are not used but it was a good exercise
package etcd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	etcdClient "go.etcd.io/etcd/client"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/pkg/transport"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type SecureCfg struct {
	Cert   string
	Key    string
	CaCert string
}

// Transport instantiate a new https Transport client
func Transport(ca string, clientCertificate string, clientCertificateKey string) *http.Transport {
	caCert, err := ioutil.ReadFile(ca)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, err := tls.LoadX509KeyPair(clientCertificate, clientCertificateKey)
	if err != nil {
		log.Fatal(err)
	}

	var t = &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      caCertPool,
			Certificates: []tls.Certificate{cert},
		},
	}
	return t
}

// Client instantiate a new etcd client, do nothing for now.
// Not totally nothing, it puts a key and read it.
func HttpClient(transport *http.Transport, endPoints string) {
	cfg := etcdClient.Config{
		Endpoints: []string{endPoints},
		Transport: transport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := etcdClient.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := etcdClient.NewKeysAPI(c)
	// set "/foo" key with "bar" value
	log.Print("Setting '/foo' key with 'bar' value")
	resp, err := kapi.Set(context.Background(), "/foo", "bar", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Set is done. Metadata is %q\n", resp)
	}
	// get "/foo" key's value
	log.Print("Getting '/foo' key value")
	//resp, Err = kapi.Get(context.Background(), "/foo", nil)
	resp, err = kapi.Get(context.Background(), "/foo", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Get is done. Metadata is %q\n", resp)
		// print value
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}
}

//ClientConfig returns an ETCD TLS ready client
func GrpcClient(tlsCfg SecureCfg, endPoints []string) *clientv3.Client {
	// populate the certificate and the key
	var transportTls *transport.TLSInfo
	tlsInfo := transport.TLSInfo{
		CertFile:            tlsCfg.Cert,
		KeyFile:             tlsCfg.Key,
		TrustedCAFile:       tlsCfg.CaCert,
		ClientCertAuth:      false,
		CRLFile:             "",
		InsecureSkipVerify:  false,
		SkipClientSANVerify: false,
		ServerName:          "",
		HandshakeFailure:    nil,
		CipherSuites:        nil,
		AllowedCN:           "",
		AllowedHostname:     "",
		Logger:              nil,
		EmptyCN:             false,
	}
	transportTls = &tlsInfo
	clientTls, err := transportTls.ClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	var cfgs = clientv3.Config{
		Endpoints:            endPoints,
		AutoSyncInterval:     0,
		DialTimeout:          1000000,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
		MaxCallSendMsgSize:   0,
		MaxCallRecvMsgSize:   0,
		TLS:                  clientTls,
		Username:             "",
		Password:             "",
		RejectOldCluster:     false,
		DialOptions:          nil,
		Context:              nil,
		LogConfig:            nil,
		PermitWithoutStream:  false,
	}
	cli, err := clientv3.New(cfgs)
	if err != nil {
		log.Fatal(err)
	}
	return cli
}
