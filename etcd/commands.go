package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
)

func ClusterStatus(cli clientv3.Client) bool {
	resp, err := cli.Status(context.Background(), cli.Endpoints()[0])
	if err != nil {
		log.Fatalln(err)
		return false
	}
	fmt.Println(resp)
	return true
}
