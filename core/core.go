// Package core contains all functions to test the cluster
package core

import (
	"github.com/stregatto/etcd_check/etcd"
	"math"
	"sort"
)

// Members check, verifies if some member is not available
func MembersHealthiness(raftIdxPerMember etcd.RaftIndexPerMember, maxFailingMembers int) (bool, []string) {
	status := true
	var failedMembersList []string
	failedMembers := 0
	for v := range raftIdxPerMember {
		if raftIdxPerMember[v].Err != nil {
			failedMembers++
			failedMembersList = append(failedMembersList, v)
			if failedMembers >= maxFailingMembers {
				status = false
			}
		}
	}
	return status, failedMembersList
}

// RaftCoherence check if the raft index for every member is in the maxRaftDrift value.
func RaftCoherence(raftIndexPerMember etcd.RaftIndexPerMember, maxRaftDrift int) (bool, RaftValue) {

	//TODO: Divides in multiple functions
	var f = map[uint64][]string{}
	for k, v := range raftIndexPerMember {
		f[v.RaftIndex] = append(f[v.RaftIndex], k)
	}

	// everything is ok, I can exit right now.
	if len(f) <= 1 {
		return true, RaftValue{}
	}

	// define the quorum, it's supposed the ETCD cluster has odd number of members
	quorum := len(raftIndexPerMember)/2 + 1

	var raftWihtMajorityOfMembers = uint64(0) // temporary max raft
	var maxNodes = 0                          // temporary max node per raft
	var raftQuorumValue = uint64(0)           // The raft quorum it exists, if not the value is 0
	for k, v := range f {
		// define the raft that has the max number of nodes
		if len(v) > maxNodes {
			maxNodes = len(v)
			raftWihtMajorityOfMembers = k
		}
		// define the raft index the quorum has, if exists.
		if len(v) >= quorum && raftQuorumValue == 0 {
			raftQuorumValue = k
		}
	}

	// compact the map using, the raft should be not outside the raft interval
	for k, v := range f {
		if IsBetween(k, raftWihtMajorityOfMembers-1, raftWihtMajorityOfMembers+1) && k != raftWihtMajorityOfMembers {
			f[raftWihtMajorityOfMembers] = append(f[raftWihtMajorityOfMembers], v...)
			delete(f, k)
		}
	}

	// TODO: verify if can be refactored in the previous functions
	// evaluate the f map of compacted rafts accordingly to quorum value, quorum = 0 means no quorum reached
	var failedMembers = RaftValue{}
	for k, v := range f {
		if (len(v) < quorum) || quorum == 0 {
			failedMembers[k] = append(failedMembers[k], v...)
		}
	}

	// define if rafts are outside the drift interval
	var collectedRafts []uint64
	for k := range f {
		collectedRafts = append(collectedRafts, k)
	}

	sort.Slice(collectedRafts, func(i, j int) bool {
		return collectedRafts[i] < collectedRafts[j]
	})

	// if between min Raft and max Raft there's more than drift cluster failed
	if math.Abs(float64(int64(collectedRafts[0]-collectedRafts[len(collectedRafts)-1]))) > float64(maxRaftDrift) {
		return false, failedMembers
	}

	return true, failedMembers

}
