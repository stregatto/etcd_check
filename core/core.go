// Package core contains all functions to test the cluster
package core

import (
	"fmt"
	"github.com/stregatto/etcd_check/etcd"
	"io/ioutil"
	"log"
	"math"
	"sort"
)

// Members check, verifies if some member is not available
func MembresHealthiness(raftIdxPerMember etcd.RaftIndexPerMember, maxFailingMembers int) (bool, []string) {
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
func RaftCoherence(raftIndexPerMember etcd.RaftIndexPerMember, maxRaftDrift int) (bool, raftValue) {

	//TODO: now I've the map of frequencies, I need to retrieve the members are failing accordingly to raft drift or return all members
	var f = map[uint64][]string{}
	for k, v := range raftIndexPerMember {
		f[v.RaftIndex] = append(f[v.RaftIndex], k)
	}

	// everything is ok, I can exit right now.
	if len(f) <= 1 {
		fmt.Println("fast exit!")
		return true, raftValue{}
	}

	//quorum := len(raftIndexPerMember)/2 + 1
	var failedMembers = raftValue{}
	collectedRafts := []uint64{}

	for k := range f {
		collectedRafts = append(collectedRafts, k)
	}

	sort.Slice(collectedRafts, func(i, j int) bool {
		return collectedRafts[i] < collectedRafts[j]
	})

	if math.Abs(float64(int64(collectedRafts[0]-collectedRafts[len(collectedRafts)-1]))) > float64(maxRaftDrift) {
		return false, failedMembers
	}

	return true, failedMembers

}

// GetFile returns a file content in []byte format form a given path, useless.
func GetFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
