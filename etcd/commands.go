package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"sync"
)

func ClusterStatus(cli *clientv3.Client) bool {
	var wg sync.WaitGroup
	hch := make(chan *clientv3.StatusResponse, len(cli.Endpoints()))
	for _, ep := range cli.Endpoints() {
		wg.Add(1)
		go func(cli *clientv3.Client) {
			defer wg.Done()
			resp, err := cli.Status(context.Background(), ep)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(resp.RaftAppliedIndex)
			fmt.Println(resp.RaftIndex)
			fmt.Println(resp.RaftTerm)
			hch <- resp
		}(cli)
	}
	wg.Wait()
	close(hch)
	return true
}
