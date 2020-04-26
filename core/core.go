// Package core contains all functions to test the cluster
package core

import (
	"github.com/stregatto/etcd_check/etcd"
	"io/ioutil"
	"log"
	"math"
	"sort"
)

// Members check, verifies if some member is not available
func MembresHealthiness(raftIdxPerMember []etcd.RaftIndexPerMember, maxFailingMembers int) (bool, []string) {
	status := true
	var failedMembersList []string
	failedMembers := 0
	for _, rip := range raftIdxPerMember {
		if rip.Err != nil {
			failedMembers++
			failedMembersList = append(failedMembersList, rip.Server)
			if failedMembers >= maxFailingMembers {
				status = false
			}
		}
	}
	return status, failedMembersList
}

// RaftCoherence check if the raft index for every member is in the maxRaftDrift value.
// TODO: fix the function accordingly to tests
func RaftCoherence(raftIdxPerMember []etcd.RaftIndexPerMember, maxRaftDrift int) (bool, []raftValue) {
	status := true
	sort.Slice(raftIdxPerMember, func(i, j int) bool {
		if raftIdxPerMember[i].RaftIndex < raftIdxPerMember[j].RaftIndex {
			return true
		}
		if raftIdxPerMember[i].RaftIndex > raftIdxPerMember[j].RaftIndex {
			return false
		}
		return raftIdxPerMember[i].RaftIndex < raftIdxPerMember[j].RaftIndex
	})

	if math.Abs(float64(raftIdxPerMember[0].RaftIndex-raftIdxPerMember[len(raftIdxPerMember)-1].RaftIndex)) > float64(maxRaftDrift) {
		status = false
	}

	//TODO: now I've the map of frequencies, I need to retrive the members are failing accordingly to raft drift or return all members
	var f = map[uint64][]string{
		raftIdxPerMember[0].RaftIndex: {},
	}
	for _, v := range raftIdxPerMember {
		f[v.RaftIndex] = append(f[v.RaftIndex], v.Server)
	}

	return status, []raftValue{
		{10,
			"etcd1"},
	}
}

// GetFile returns a file content in []byte format form a given path, useless.
func GetFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
