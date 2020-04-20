//Package etcd contains all tools to manage etcd client and query an etcd cluster
package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"sync"
)

//command provides the interface to query all cluster member sending to it the command you need
func command(cli *clientv3.Client, cmd string) <-chan commandResChan {
	var wg sync.WaitGroup
	// channel still not ok
	// hch := make(chan commandResChan, len(cli.Endpoints()))
	hch := make(chan commandResChan)
	for _, ep := range cli.Endpoints() {
		wg.Add(1)
		go func(cli *clientv3.Client, ep string) {
			var crc commandResChan
			defer wg.Done()
			crc.res = callCommand(cli, cmd, ep)
			crc.server = ep
			hch <- crc
		}(cli, ep)
	}
	go func() {
		wg.Wait()
		close(hch)
	}()
	return hch
}

// callCommand execute the command cmd against one ETCD endpoint,
// it returns the interface{} containing the result of the command returned by clientv3
func callCommand(cli *clientv3.Client, cmd, ep string) interface{} {
	switch cmd {
	case "clusterStatus":
		var resp, err = cli.Status(context.Background(), ep)
		if err != nil {
			log.Fatal(err)
		}
		return resp
	case "endpoints":
		var resp = cli.Endpoints()
		return resp
	case "memberList":
		var resp, err = cli.MemberList(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		return resp
	}
	return fmt.Errorf("missing Command")
}

//GetRaftIndexPerMembers returns a list of raftIndexPerMember
func GetRaftIndexPerMembers(cli *clientv3.Client) []raftIndexPerMember {
	ch := command(cli, "clusterStatus")
	raftx := make([]raftIndexPerMember, len(ch))
	for v := range ch {
		raftx = append(raftx, raftIndexPerMember{
			v.server,
			v.res.(*clientv3.StatusResponse).Header.MemberId,
			v.res.(*clientv3.StatusResponse).RaftIndex,
		})

	}
	return raftx
}

//TODO
//fix the strunct that retrives the ep from the resault... no idea

//GetEndPointsFromInitiatedClient returns list of endpoints form the client you instantiated.
func GetEndPointsFromInitiatedClient(cli *clientv3.Client) []clientEndpoints {
	ch := command(cli, "endpoints")
	epx := make([]clientEndpoints, len(ch))
	for v := range ch {
		epx = append(epx, clientEndpoints{
			v.server,
			v.res.([]string)})
	}
	return epx
}

//GetClusterEndpoints returns the list of endpoint configured in the cluster you are querying.
func GetClusterEndpoints(cli *clientv3.Client) []clusterEndpoints {
	ch := command(cli, "memberList")
	epx := make([]clusterEndpoints, len(ch))
	for v := range ch {
		epx = append(epx, clusterEndpoints{v.server, v.res.(*clientv3.MemberListResponse).Members})
	}
	return epx
}
