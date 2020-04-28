//Package etcd contains all tools to manage etcd client and query an etcd cluster
package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"sync"
	"time"
)

//command provides the interface to query all cluster member sending to it the command you need
func command(cli *clientv3.Client, cmd string) <-chan commandResChan {
	var wg sync.WaitGroup
	hch := make(chan commandResChan)
	for _, ep := range cli.Endpoints() {
		wg.Add(1)
		go func(cli *clientv3.Client, ep string) {
			var crc commandResChan
			defer wg.Done()
			crc.res, crc.err = callCommand(cli, cmd, ep)
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
func callCommand(cli *clientv3.Client, cmd, ep string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	switch cmd {
	case "clusterStatus":
		var res, err = cli.Status(ctx, ep)
		return res, err
	case "endpoints":
		var res = cli.Endpoints()
		return res, nil
	case "memberList":
		var res, err = cli.MemberList(ctx)
		return res, err
	}
	return nil, fmt.Errorf("missing Command")
}

//GetRaftIndexPerMembers returns a list of RaftIndexPerMember
func GetRaftIndexPerMembers(cli *clientv3.Client) RaftIndexPerMember {
	ch := command(cli, "clusterStatus")
	raftx := make(RaftIndexPerMember, len(ch))
	for v := range ch {
		raftx[v.server] = RaftData{
			MemberId: func() uint64 {
				if v.err == nil {
					return v.res.(*clientv3.StatusResponse).Header.MemberId
				}
				return uint64(0)
			}(),
			RaftIndex: func() uint64 {
				if v.err == nil {
					return v.res.(*clientv3.StatusResponse).RaftIndex
				}
				return uint64(0)
			}(),
			Err: v.err,
		}

	}
	return raftx
}

//GetEndPointsFromInitiatedClient returns list of endpoints form the client you instantiated.
func GetEndPointsFromInitiatedClient(cli *clientv3.Client) []ClientEndpoints {
	ch := command(cli, "endpoints")
	epx := make([]ClientEndpoints, len(ch))
	for v := range ch {
		epx = append(epx, ClientEndpoints{
			v.server,
			v.res.([]string),
			v.err,
		})
	}
	return epx
}

//GetClusterEndpoints returns the list of endpoint configured in the cluster you are querying.
func GetClusterEndpoints(cli *clientv3.Client) []ClusterEndpoints {
	ch := command(cli, "memberList")
	epx := make([]ClusterEndpoints, len(ch))
	for v := range ch {
		epx = append(epx, ClusterEndpoints{
			v.server,
			v.res.(*clientv3.MemberListResponse).Members,
			v.err,
		})
	}
	return epx
}
